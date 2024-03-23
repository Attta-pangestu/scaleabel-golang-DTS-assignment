package controllers

import (
	models "MyGramAtta/models"
	repo "MyGramAtta/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary Post Photo
// @Description Post a new Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param PostPhoto body models.RequestPhoto true "Post photo"
// @Success 201 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Router /photo/post [post]
func CreatePhoto(ctx *gin.Context) {
	db := repo.GetDB()
	fmt.Println("Running CreatePhoto")

	userData, ok := ctx.Get("userData")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Data pengguna tidak ditemukan",
		})
		return
	}

	claims, ok := userData.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Data pengguna tidak valid",
		})
		return
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "ID pengguna tidak valid",
		})
		return
	}

	var user models.User
	errUser := db.First(&user, uint(userID)).Error
	if errUser != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	
	fmt.Println("userId: ", userID) 

	var requestPhoto models.RequestPhoto
	err := ctx.ShouldBind(&requestPhoto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	

	photo := models.Photo{
		UserID:   uint(userID),
		PhotoUrl: requestPhoto.PhotoUrl,
		Title:    requestPhoto.Title,
		Caption:  requestPhoto.Caption,
	}

	err = db.Debug().Create(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, response)
}


// GetPhotos godoc
// @Summary Dapatkan detail semua foto
// @Description Dapatkan detail semua foto atau tambahkan parameter kueri user_id untuk semua foto dari user_id (opsional)
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query integer false "Dapatkan semua foto yang difilter berdasarkan user_id"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photo/getAll [get]
	
func GetAllPhotos(ctx *gin.Context) {
	var photos []models.Photo

	db := repo.GetDB()

	userData, ok := ctx.Get("userData")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Data pengguna tidak ditemukan",
		})
		return
	}

	claims, ok := userData.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Data pengguna tidak valid",
		})
		return
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "ID pengguna tidak valid",
		})
		return
	}

	var err error
	if userIDQuery, ok := ctx.GetQuery("user_id"); ok {
		userID, err = strconv.ParseFloat(userIDQuery, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id harus berupa angka",
			})
			return
		}
	}

	err = db.Debug().Order("id").Where("user_id = ?", uint(userID)).Find(&photos).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var responsePhotos []models.PhotoResponse
	for _, photo := range photos {
		responsePhotos = append(responsePhotos, models.PhotoResponse{
			 ID:         photo.ID,
			Title:      photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:   photo.PhotoUrl,
			UserID:     photo.UserID,
			CreatedAt:  photo.CreatedAt,
			UpdatedAt:  photo.UpdatedAt,
			User: models.UserResponse{
				Email:    userData.(jwt.MapClaims)["email"].(string),
				Username: userData.(jwt.MapClaims)["username"].(string),
			},
		})
	}

	// Mengirimkan respons dengan data foto yang sudah dimodifikasi
	ctx.JSON(http.StatusOK, responsePhotos)
}


// UpdatePhoto godoc
// @Summary Updated data photo with socialMediaID
// @Description Update data photo by id, NOTE: photo is not updated, just title and caption can be updated, so in the body photo_url doesn't use
// @Tags Photo
// @Accept json
// @Produce json
// @Param photoID path integer true "photoID of the data photo to be updated"
// @Param UpdatePhoto body models.RequestPhoto true "Update photo"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photo/update/{photoID} [put]
func UpdatePhoto(ctx *gin.Context) {
	var photo, findPhoto models.Photo

	db := repo.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&findPhoto).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Photo not found",
		})
		return
	}

	_ = ctx.ShouldBind(&photo)

	photo = models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: findPhoto.PhotoUrl,
	}

	photo.ID = uint(photoID)
	photo.CreatedAt = findPhoto.CreatedAt
	photo.UserID = findPhoto.UserID

	err = db.Debug().Model(&photo).Where("id = ?", photoID).Updates(photo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary Delete data photo
// @Description Delete data photo by id
// @Tags Photo
// @Accept json
// @Produce json
// @Security
// @Param photoID path integer true "photoID of the data photo to be deleted"
// @Success 200 {object} models.Photo
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /photo/delete/{photoID} [delete]
func DeletePhoto(ctx *gin.Context) {
	var photo models.Photo

	db := repo.GetDB()

	photoID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Where("id = ?", photoID).First(&photo).Delete(&photo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "Your photo has been successfully deleted",
	}
	ctx.JSON(http.StatusOK,response)
}
