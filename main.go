package main

import (
	"github.com/gin-gonic/gin"
	"GoBoard/article"
)


func main() {
	router := gin.Default()
	article.AddRouter(router)
	//user.AddRouter(router)
	router.Run()
}