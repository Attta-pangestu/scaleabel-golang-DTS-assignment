package controllers

import (
	helpers "MyGramAtta/helper"
	models "MyGramAtta/models"
	repo "MyGramAtta/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
// @Summary Membuat media sosial
// @Description Membuat media sosial baru untuk pengguna yang terautentikasi
// @Tags Media Sosial
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Token akses JWT"
// @Param name body string true "Nama media sosial"
// @Param social_media_url body string true "URL media sosial"
// @Success 201 {object} models.SocialMedia "Media sosial berhasil dibuat"
// @Failure 400 {object} models.ResponseFailed "Permintaan tidak valid"
// @Failure 401 {object} models.ResponseFailedUnauthorized "Tidak terautentikasi"
// @Router /social-media/create [post]
func CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":            socialMedia.ID,
		"name":          socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":       socialMedia.UserID,
		"created_at":    socialMedia.CreatedAt,
	})
}

// GetAllSocialMedia godoc
// @Summary Get details of all social media
// @Description Get details of all social media or add query parameter user_id for all social media from user_id (optional)
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "Get all social media filter by user_id"
// @Success 200 {object} gin.H{"social_medias": []models.SocialMedia}
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/getAll [get]
func GetAllSocialMedia(ctx *gin.Context) {
	var socialMedia []models.SocialMedia

	db := repo.GetDB()

	if _, ok := ctx.GetQuery("user_id"); ok {
		userID, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Input user_id dengan angka",
			})
			return
		}

		err = db.Debug().Preload("User").Order("id").Where("user_id = ?", userID).Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if len(socialMedia) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("user_id %d tidak memiliki media sosial", userID),
			})
			return
		}
	} else {
		err := db.Debug().Preload("User").Order("id").Find(&socialMedia).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	var response []gin.H
	for _, s := range socialMedia {
		res := gin.H{
			"id":               s.ID,
			"created_at":       s.CreatedAt,
			"updated_at":       s.UpdatedAt,
			"name":             s.Name,
			"social_media_url": s.SocialMediaUrl,
			"user_id":          s.UserID,
			"User": gin.H{
				"id":       s.User.ID,
				"username": s.User.Username,
				"email":    s.User.Email,
			},
		}
		response = append(response, res)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": response,
	})
}


// UpdateSocialMedia godoc
// @Summary Updated data social media
// @Description Update data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param socialMediaID path integer true "socialMediaID of the data social media to be updated"
// @Param SocialMedia body models.RequestSocialMedia true "updated social media"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/update/{socialMediaID} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	var socialMedia, findSocialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&socialMedia); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err = db.Debug().Where("id = ?", socialMediaID).First(&findSocialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Social media not found",
		})
		return
	}

	socialMedia = models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	socialMedia.ID = uint(socialMediaID)
	socialMedia.CreatedAt = findSocialMedia.CreatedAt
	socialMedia.UserID = findSocialMedia.UserID

	err = db.Debug().Model(&socialMedia).Where("id = ?", socialMediaID).Updates(socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete data social media
// @Description Delete data social media by id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param socialMediaID path integer true "socialMediaID of the data social media to be deleted"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /social-media/delete/{socialMediaID} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	db := repo.GetDB()

	err = db.Debug().Where("id = ?", socialMediaID).First(&socialMedia).Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Social Media %s successfully deleted", socialMedia.Name),
	})
}
