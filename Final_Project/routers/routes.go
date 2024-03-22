package routers

import (
	controllers "MyGramAtta/controllers"
	middlewares "MyGramAtta/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		
		userRouter.POST("/register", controllers.RegisterUser)
		
		userRouter.POST("/login", controllers.LoginUser)
		
		userRouter.Use(middlewares.UserAuthentication())
		userRouter.PUT("/", middlewares.UserAuthentication(), controllers.UpdateUser)
		userRouter.DELETE("/", middlewares.UserAuthentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.UserAuthentication())

		photoRouter.POST("/", controllers.CreatePhoto)

		photoRouter.GET("/", controllers.GetAllPhotos)

		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.UserAuthentication())

		commentRouter.POST("/", controllers.CreateComment)

		commentRouter.GET("/", controllers.GetAllComments)


		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)

		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComent)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.UserAuthentication())

		socialMediaRouter.POST("/", controllers.CreateSocialMedia)

		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)

		socialMediaRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

		socialMediaRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}