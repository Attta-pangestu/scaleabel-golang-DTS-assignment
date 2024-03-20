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
	}

	r.Static("/img", "./assets")
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authetication())

		photoRouter.POST("/post", controllers.CreatePhoto)

		photoRouter.GET("/getAll", controllers.GetAllPhoto)

		photoRouter.GET("/getOne/:photoID", controllers.GetOnePhoto)

		photoRouter.PUT("/update/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		photoRouter.DELETE("/delete/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authetication())

		commentRouter.POST("/create", controllers.CreateComment)

		commentRouter.GET("/getAll", controllers.GetAllComent)

		commentRouter.GET("/getOne/:commentID", controllers.GetOneComment)

		commentRouter.PUT("/update/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)

		commentRouter.DELETE("/delete/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComent)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authetication())

		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)

		socialMediaRouter.GET("/getAll", controllers.GetAllSocialMedia)

		socialMediaRouter.GET("/getOne/:socialMediaID", controllers.GetOneSocialMedia)

		socialMediaRouter.PUT("/update/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

		socialMediaRouter.DELETE("/delete/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}