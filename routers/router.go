package routers

import (
	"github.com/gin-gonic/gin"
	"pjt_article/Controllers"
)

func InitRouter() {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/", Controllers.CreateUser)
		user.GET("/", Controllers.GetAllUsers)
		//	user.GET("/:id", getUser)
		//	user.DELETE("/:id", deleteUser)
	}

	article := router.Group("/article")
	{
		article.POST("/", Controllers.CreateArticle)
		article.GET("/", Controllers.GetAllArticles)
		article.GET("/:id", Controllers.GetArticle)
		article.PUT("/:id", Controllers.UpdateArticle)
		article.DELETE("/:id", Controllers.DeleteArticle)


	}

	test := router.Group("/test")
	{
		test.GET("/:id", Controllers.GetArticleListByUserId)
	}

	router.Run(":8080")
}
