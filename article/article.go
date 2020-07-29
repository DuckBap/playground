package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"jwon/database"
	"jwon/structInfo"
)

func CreateArticle(c *gin.Context) {
	var I structInfo.Article
	title := c.PostForm("Title")
	body := c.PostForm("Body")
	//if title == "" || body == "" {
	//	fmt.Println("Fields can't be black")
	//} else {
	//	I.Title = title
	//	I.Body = body
	//	//I.Title = c.PostForm("Title")
	//	//I.Body = c.PostForm("Body")
	//	Db.Create(&I)
	//	c.JSON(200, I)
	//}
	I.Title = title
	I.Body = body
	database.Db.Create(&I)
	c.JSON(200, I)
}

func ListArticles(c *gin.Context) {
	var I []structInfo.Article
	if err := database.Db.Find(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, I)
	}
}

func ViewArticle(c *gin.Context) {
	var I structInfo.Article
	id := c.Params.ByName("id")
	if err := database.Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, I)
	}
}

func UpdateArticle(c *gin.Context) {
	var I structInfo.Article
	id := c.Params.ByName("id")
	//title := c.PostForm("Title")
	//body := c.PostForm("Body")
	if err := database.Db.Where("id = ?", id).First(&I).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		//I.Title = title
		//I.Body = body
		I.Title = c.PostForm("Title")
		I.Body = c.PostForm("Body")
		database.Db.Save(&I)
		c.JSON(200, I)
	}
}

func DeleteArticle(c *gin.Context) {
	var I structInfo.Article
	id := c.Params.ByName("id")
	//Db.Delete(&I)
	msg := database.Db.Where("id = ? ", id).Delete(&I)
	fmt.Println(msg)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func HandleArticle(router *gin.Engine) {
	router.POST("/article", CreateArticle)
	router.GET("/article", ListArticles)
	router.GET("/article/:id", ViewArticle)
	router.PUT("/article/:id", UpdateArticle)
	router.DELETE("/article/:id", DeleteArticle)
}