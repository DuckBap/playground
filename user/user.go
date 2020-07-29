package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"jwon/database"
	"jwon/structInfo"
)

func CreateUser(c *gin.Context) {
	var I structInfo.User
	//username := c.PostForm("Username")
	//password := c.PostForm("Password")
	//nickname := c.PostForm("Nickname")
	//I.Username = username
	//I.Password = password
	//I.Nickname = nickname
	I.Username = c.PostForm("Username")
	I.Password = c.PostForm("Password")
	I.Nickname = c.PostForm("Nickname")
	database.Db.Create(&I)
	c.JSON(200, I)
}

func ListUsers(c *gin.Context) {
	var I []structInfo.User
	if err := database.Db.Find(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, I)
	}
}

func GetUser(c *gin.Context) {
	var I structInfo.User
	id := c.Params.ByName("id")
	if err := database.Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, I)
	}
}

func UpdateUser(c *gin.Context) {
	var I structInfo.User
	id := c.Params.ByName("id")
	//username := c.PostForm("Username")
	//password := c.PostForm("Password")
	//nickname := c.PostForm("Nickname")
	if err := database.Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		//I.Username = username
		//I.Password = password
		//I.Nickname = nickname
		I.Username = c.PostForm("Username")
		I.Password = c.PostForm("Password")
		I.Nickname = c.PostForm("Nickname")
		database.Db.Save(&I)
		c.JSON(200, I)
	}
}

func DeleteUser(c *gin.Context) {
	var I structInfo.User
	id := c.Params.ByName("id")
	msg := database.Db.Where("id = ? ", id).Delete(&I)
	fmt.Println(msg)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func HandleUser(router *gin.Engine) {
	router.POST("/user", CreateUser)
	router.GET("/user", ListUsers)
	router.GET("/user/:id", GetUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
}