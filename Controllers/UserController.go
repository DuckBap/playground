package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pjt_article/Models"
	"pjt_article/Repositories"
)

func CreateUser(c *gin.Context) {
/*
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	user := Models.User{Username: username, Password: password, Nickname: nickname}
*/
	var user Models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := Repositories.CreateUser(&user); err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    &user,
		"message": "created",
	})
}

func GetAllUsers(c *gin.Context){
	var users []Models.User
	err := Repositories.GetAllUsers(&users)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(c *gin.Context) {
	var article Models.Article
	id := c.Param("id")
	err := Repositories.GetArticle(&article, id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, article)
	}
}

func UpdateUser(c *gin.Context) {
	var article Models.Article
	id := c.Param("id")
	err := Repositories.GetArticle(&article, id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		body := c.PostForm("body")
		err = Repositories.UpdateArticle(&article, body)

		if err != nil {

		} else {
			c.JSON(http.StatusOK, article)
		}
	}
}

func DeleteUser(c *gin.Context) {
	var article Models.Article
	id := c.Param("id")
	err := Repositories.GetArticle(&article, id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		err = Repositories.DeleteArticle(&article)
		if err != nil {

		} else {
			c.JSON(200, "deleted")
		}
	}
}
