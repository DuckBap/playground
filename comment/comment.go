package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwon/structInfo"
)

var Db *gorm.DB
var Err error


func CreateComment(c *gin.Context) {
	var I structInfo.Comment
	body := c.PostForm("Body")
	I.Body = body
	//c.BindJSON(&I)
	Db.Create(&I)
	c.JSON(200, I)
}

func ViewComment(c *gin.Context) {
	var I structInfo.Comment
	id := c.Params.ByName("id")
	if err := Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, I)
	}
}

func UpdateComment(c *gin.Context) {
	var I structInfo.Comment
	id := c.Params.ByName("id")
	body := c.PostForm("Body")
	if err := Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		I.Body = body
		Db.Save(&I)
		c.JSON(200, I)
	}
}

func DeleteComment(c *gin.Context) {
	var I structInfo.Comment
	id := c.Params.ByName("id")
	msg := Db.Where("id = ? ", id).Delete(&I)
	fmt.Println(msg)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func ConfigArticle(router *gin.Engine) {
	router.POST("/comment", CreateComment)
	router.GET("/comment/:id", ViewComment)
	router.PUT("/comment/:id", UpdateComment)
	router.DELETE("/comment/:id", DeleteComment)
}