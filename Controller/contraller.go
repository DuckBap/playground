package Controller

import (
	"example.com/test/Model"
	"github.com/gin-gonic/gin"
)

func UrlController() {
	router := gin.Default()

	//router.POST("/:param", Model.ParesData)
	//router.POST("/user", Model.ParseUserData)
	//router.POST("/article", Model.ParseArticleData)
	router.POST("/:param", Model.ParseData)
	//router.POST("/mi", Model.Migrate)
	//router.POST("/comment", Model.ParseCommentData)
	//router.POST("/test", test)
	//router.GET("/comment", Model.ReadCommentData)
	router.GET("/user", Model.ReadUserData)
	router.GET("/article", Model.ReadArticleData)
	router.Run(":8080")
}
//
//func test(c *gin.Context) {
//	fmt.Println(c.MultipartForm())
//}