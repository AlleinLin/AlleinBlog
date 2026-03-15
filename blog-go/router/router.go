package router

import (
	"blog-go/controller"
	"blog-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.Recovery())

	userController := controller.NewUserController()
	articleController := controller.NewArticleController()
	categoryController := controller.NewCategoryController()
	tagController := controller.NewTagController()
	commentController := controller.NewCommentController()

	r.POST("/login", userController.Login)
	r.POST("/register", userController.Register)
	r.GET("/user/admin", userController.GetAdminInfo)

	auth := r.Group("")
	auth.Use(middleware.JWTAuth())
	{
		auth.POST("/logout", userController.Logout)
		auth.GET("/user/info", userController.GetUserInfo)
		auth.PUT("/user/info", userController.UpdateUserInfo)

		auth.POST("/article", articleController.AddArticle)
		auth.PUT("/article", articleController.UpdateArticle)
		auth.DELETE("/article/:id", articleController.DeleteArticle)

		auth.POST("/comment", commentController.AddComment)
		auth.PUT("/comment", commentController.UpdateComment)
		auth.DELETE("/comment/:id", commentController.DeleteComment)
	}

	r.GET("/article/list", articleController.GetArticleList)
	r.GET("/article/hot", articleController.GetHotArticleList)
	r.GET("/article/:id", articleController.GetArticleDetail)
	r.GET("/article/count", articleController.GetArticleCount)
	r.GET("/article/archive", articleController.GetArchiveList)

	r.GET("/category/list", categoryController.GetCategoryList)
	r.GET("/category/:id", categoryController.GetCategoryByID)

	r.GET("/tag/list", tagController.GetTagList)
	r.GET("/tag/:id", tagController.GetTagByID)

	r.GET("/comment/list", commentController.GetCommentList)

	return r
}
