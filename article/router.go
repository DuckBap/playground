package article

import "github.com/gin-gonic/gin"

func AddRouter(router *gin.Engine) {
	appRouter := router.Group("/article")
	{
		appRouter.GET("/", listArticles)
		appRouter.POST("/", createArticle)
		appRouter.PATCH("/:id", updateArticle)
		appRouter.DELETE("/:id", deleteArticle)
		appRouter.GET("/:id", getArticle)
		appRouter.POST("/:id/comment", createComment)
	}
}