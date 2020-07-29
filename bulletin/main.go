package main

import (
	"Model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	Model.Db, Model.Err = gorm.Open(mysql.Open("yechoi:@/bulletin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if Model.Err != nil {
		panic("failed to connect database")
	}
	Model.Db.AutoMigrate(&Model.User{}, &Model.Article{})

	//defer Article.Db.Close()
	//defer User.Db.Close()

	router := gin.Default()
	Model.EditUser(router)
	Model.EditArticle(router)
	router.Run(":8080")
}