package main

import (
	"github.com/gin-gonic/gin"
	"jwon/article"
	//"jwon/database"
	"jwon/user"
)

func main() {
	//database.init()
	//defer user.Db.Close()
	//defer user.Db.Close()
	//defer article.Db.Close()
	router := gin.Default()
	user.HandleUser(router)
	article.HandleArticle(router)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}