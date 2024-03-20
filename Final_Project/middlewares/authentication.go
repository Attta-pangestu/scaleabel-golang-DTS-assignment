package middlewares

import (
	helpers "MyGramAtta/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authetication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyToken)

		ctx.Next()
	}
}
