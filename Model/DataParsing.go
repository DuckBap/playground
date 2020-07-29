package Model

import (
	"example.com/test/Data"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)
//type Comment struct {
//	gorm.Model
//	Body string
//	ArticleID uint
//	UserID uint
//}

//type User struct {
//	gorm.Model
//	UserName string
//	Password string
//	NickName string
//	Articles []Article
//	//Comment []Comment
//}
//
//type Article struct {
//	gorm.Model
//	Title string
//	Body string
//	UserID uint
//	//Comment []Comment
//}
//type	Article Data.Article
//type	User	Data.User
var 	dsn		string
var		db		*gorm.DB
var		err		error

func init() {
	dsn = "dokang:1234@tcp(127.0.0.1:3306)/exdb?charset=utf8mb4&parseTime=True&loc=Local"
}

func ParseData (c *gin.Context) {
	param := c.Param("param")
	fmt.Println(param)
	if param == "article" {
		ParseArticleData(c)
	} else if param == "user" {
		ParseUserData(c)
	} else {
		Migrate(c)
	}
}

//func ParseCommentData (c *gin.Context) {
//	var comment	Data.Comment
//
//	comment.Body = c.PostForm("comment_body")
//	articleID,_ := strconv.Atoi(c.PostForm("article_id"))
//	userID,_ := strconv.Atoi(c.PostForm("user_id"))
//	comment.ArticleID = uint(articleID)
//	comment.UserID = uint(userID)
//	InputDataBase(&comment, "comment")
//	c.JSON(http.StatusOK, comment)
//}

func ParseArticleData (c *gin.Context) {
	var	article	Data.Article

	//article.Title = c.PostForm("title")
	//article.Body = c.PostForm("body")
	//userID,_ := strconv.Atoi(c.PostForm("user_id"))
	article.UserID = 1
	c.BindJSON(&article)
	fmt.Println(article)
	InputDataBase(&article)
	c.JSON(http.StatusOK, article)
}

func ParseUserData (c *gin.Context) {
	var	user	Data.User

	//user.NickName = c.PostForm("nickname")
	//user.UserName = c.PostForm("username")
	//user.Password = c.PostForm("password")
	c.ShouldBindJSON(&user)
	InputDataBase(&user)
	c.JSON(http.StatusOK, user)
}

func InputDataBase (product interface{}) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	//db.AutoMigrate()
	mysqlDB, err := db.DB()
	if err != nil {
		panic("error")
	}
	defer mysqlDB.Close()
	db.AutoMigrate(product)
	db.Create(product)
}

func Migrate(c *gin.Context) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	//db.AutoMigrate()
	db.AutoMigrate(&Data.User{}, &Data.Article{})
}
