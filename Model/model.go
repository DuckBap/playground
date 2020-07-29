package Model

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	_"gorm.io/driver/mysql"
	"strconv"
)

var Db *gorm.DB
var Err error

type Article struct {
	gorm.Model
	Title string
	Body string
	UserID uint
	// Comments []Comment
}

type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
	Articles []Article
	// Comments []Comment
}

func EditUser(router *gin.Engine) {
	router.POST("/user", CreateUser)
	router.GET("/user", GetsUser)
	router.GET("/user/:id", GetUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
}

func CreateUser(c *gin.Context){
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	Nickname := c.PostForm("Nickname")

	var user User
	user.Username = Username
	user.Password = Password
	user.Nickname = Nickname
	Db.Model(&user).Association("Articles")
	Db.Create(&user)

	c.JSON(200, gin.H{
		"Username" : Username,
		"Password" : Password,
		"Nickname" : Nickname,
	})
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if Err := Db.Where("id = ?", id).First(&user).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		c.JSON(200, user)
	}
}

func GetsUser(c *gin.Context) {
	var users []User
	if Err := Db.Find(&users).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		c.JSON(200, users)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	Username := c.PostForm("Username")
	Password := c.PostForm("Password")
	Nickname := c.PostForm("Nickname")

	var user User
	if Err := Db.Where("id = ?", id).First(&user).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		user.Username = Username
		user.Password = Password
		user.Nickname = Nickname
		Db.Save(&user)
		c.JSON(200, user)
	}
}

func DeleteUser(c *gin.Context){
	id := c.Params.ByName("id")
	var user User
	d := Db.Where("id = ? ", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func EditArticle(router *gin.Engine) {
	router.POST("/article", CreateArticle)
	router.GET("/article", GetsArticle)
	router.GET("/article/:id", GetArticle)
	router.PUT("/article/:id", UpdateArticle)
	router.DELETE("/article/:id", DeleteArticle)
}

func CreateArticle(c *gin.Context){
	Title := c.PostForm("Title")
	Body := c.PostForm("Body")
	UserID,_:= strconv.Atoi(c.PostForm("UserID"))

	var article Article
	article.Title = Title
	article.Body = Body
	article.UserID = uint(UserID)
	Db.Create(&article)

	c.JSON(200, gin.H{
		"Title" : Title,
		"Body" : Body,
		"UserID" : UserID,
	})
}

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	if Err := Db.Where("id = ?", id).First(&article).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		c.JSON(200, article)
	}
}

func GetsArticle(c *gin.Context) {
	var articles []Article
	if Err := Db.Find(&articles).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		c.JSON(200, articles)
	}
}

func UpdateArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	Title := c.PostForm("Title")
	Body := c.PostForm("Body")

	var article Article
	if Err := Db.Where("id = ?", id).First(&article).Error; Err != nil {
		c.AbortWithStatus(404)
		fmt.Println(Err)
	} else {
		article.Title = Title
		article.Body = Body
		Db.Save(&article)
		c.JSON(200, article)
	}
}

func DeleteArticle(c *gin.Context){
	id := c.Params.ByName("id")
	var article Article
	d := Db.Where("id = ? ", id).Delete(&article)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}