package controllers

import (
	helpers "MyGramAtta/helper"
	models "MyGramAtta/models"
	"MyGramAtta/repo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateComment godoc
// @Summary Create Comment
// @Description Post a new Comment and add query parameter photo_id for comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photo_id query integer true "Photo for comment"
// @Param CreateComment body models.RequestComment true "Create comment"
// @Success 201 {object} models.Comment
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /comment/create [post]

func CreateComment(ctx *gin.Context) {
    var requestComment models.RequestComment
    if err := ctx.ShouldBindJSON(&requestComment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if requestComment.Message == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Komentar harus memiliki isi"})
        return
    }
    
    userData, exists := ctx.Get("userData")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Data pengguna tidak ditemukan"})
		return
	}

	userClaims, ok := userData.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Data pengguna tidak valid"})
		return
	}
    
    userID, ok := userClaims["id"].(float64)
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ID pengguna tidak valid"})
        return
    }
    
    photoID := requestComment.PhotoID

    if photoID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "photo_id diperlukan"})
        return
    }
    

    db := repo.GetDB()

    var photo models.Photo
    if err := db.First(&photo, "id = ?", photoID).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
        return
    }
    
    comment := models.Comment{
        Message: requestComment.Message,
        PhotoID: uint(photoID),
        UserID:  uint(userID),
		
    }
    
    if err := db.Create(&comment).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    response := gin.H{
        "id":         comment.ID,
        "message":    comment.Message,
        "photo_id":   photoID,
        "user_id":    uint(userID),
        "created_at": comment.CreatedAt,
    }

    ctx.JSON(http.StatusCreated,gin.H{
		"data" : response,
	} )
}




// GetAllComment godoc
// @Summary Get details of All comment
// @Description Get details of all comment or add query parameter photo_id for all comment from photo_id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param photo_id query integer false "Get all comment from photo_id"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /comment/getAll [get]
func GetAllComments(ctx *gin.Context) {
    var comments []models.Comment

    db := repo.GetDB()

    if err := db.Preload("Photo").Preload("User").Find(&comments).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var responseComments []gin.H
    for _, comment := range comments {
        responseComment := gin.H{
            "id":         comment.ID,
            "message":    comment.Message,
            "photo_id":   comment.PhotoID,
            "user_id":    comment.UserID,
            "created_at": comment.CreatedAt,
            "updated_at": comment.UpdatedAt,
            "User": gin.H{
                "id":       comment.User.ID,
                "email":    comment.User.Email,
                "username": comment.User.Username,
            },
            "Photo": gin.H{
                "id":        comment.Photo.ID,
                "title":     comment.Photo.Title,
                "caption":   comment.Photo.Caption,
                "photo_url": comment.Photo.PhotoUrl,
                "user_id":   comment.Photo.UserID,
            },
        }
        responseComments = append(responseComments, responseComment)
    }

    ctx.JSON(http.StatusOK, responseComments)
}




// UpdateComment godoc
// @Summary Updated data comment with commentID
// @Description Update data comment by id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param commentID path integer true "commentID of the data comment to be updated"
// @Param UpdatedComment body models.RequestComment true "Updated comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /comment/update/{commentID} [put]
func UpdateComment(ctx *gin.Context) {
	var comment models.Comment
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter id tidak valid",
		})
		return
	}
	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)
	if contentType == appJson {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	err = db.Debug().Preload("Photo").Preload("User").Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Komentar dengan ID %d tidak ditemukan", commentID),
		})
		return
	}

	err = db.Debug().Model(&comment).Updates(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Permintaan tidak valid",
			"error":   err.Error(),
		})
		return
	}

	response := gin.H{
		"id":         comment.ID,
		"title":      comment.Photo.Title,
		"caption":    comment.Photo.Caption,
		"photo_url":  comment.Photo.PhotoUrl,
		"message":    comment.Message,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}


// DeleteComment godoc
// @Summary Delete data comment with commentID
// @Description Delete data comment by id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param commentID path integer true "commentID of the data comment to be deleted"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailedUnauthorized
// @Failure 404 {object} models.ResponseFailed
// @Router /comment/delete/{commentID} [delete]
func DeleteComent(ctx *gin.Context) {
	var comment models.Comment

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Input parameter with id",
		})
		return
	}

	db := repo.GetDB()

	err = db.Debug().Where("id = ?", commentID).First(&comment).Delete(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Comment with id %d not found", commentID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

