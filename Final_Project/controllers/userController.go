package controllers

import (
	helpers "MyGramAtta/helper"
	models "MyGramAtta/models"
	repo "MyGramAtta/repo"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var appJson = "application/json"

// RegisterUser godoc
// @Summary Register User
// @Description Register user for my gram
// @Tags User
// @Accept json
// @Produce json
// @Param UserRegister body models.RequestUserRegister true "User Register"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ResponseFailed
// @Router /user/register [post]
func RegisterUser(ctx *gin.Context) {
	var user models.User

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"age":      user.Age,
		"email":    user.Email,
		"username": user.Username,
	})
}

// LoginUser godoc
// @Summary Login User
// @Description Login user for have token (jwt)
// @Tags User
// @Accept json
// @Produce json
// @Param UserLogin body models.RequestUserLogin true "User Login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailed
// @Router /user/login [post]


func LoginUser(ctx *gin.Context) {
	var user models.User

	db := repo.GetDB()

	contentType := helpers.GetHeader(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token, 
	})
}


// UpdateUser godoc
// @Summary Update User
// @Description Update user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param UserUpdate body models.RequestUserUpdate true "User Update"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailed
// @Router /user/{id} [put]


type RequestUserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UpdateUser(ctx *gin.Context) {
	var updateUser models.RequestUserUpdate
	

	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if updateUser.Email == "" || updateUser.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Email dan username harus disediakan",
		})
		return
	}

	// Perbarui data pengguna
	var user models.User
	err := db.Debug().Where("id = ?", userID).First(&user).Updates(map[string]interface{}{
		"email":    updateUser.Email,
		"username": updateUser.Username,
	}).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Gagal memperbarui data pengguna",
		})
		return
	}

	output := gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"username":  user.Username,
		"age":       user.Age,
		"updated_at": user.UpdatedAt,

	}

	ctx.JSON(http.StatusOK, output)
}


// DeleteUser godoc
// @Summary Delete User
// @Description Delete user by id
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {string} string "User successfully deleted"
// @Failure 400 {object} models.ResponseFailed
// @Failure 401 {object} models.ResponseFailed
// @Failure 404 {object} models.ResponseFailed
// @Router /user/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	db := repo.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var user models.User

	// Cek apakah pengguna ditemukan
	err := db.Debug().Where("id = ?", userID).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Pengguna tidak ditemukan",
		})
		return
	}

	// Hapus pengguna
	err = db.Debug().Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Gagal menghapus pengguna",
		})
		return
	}

	ctx.
	JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}



