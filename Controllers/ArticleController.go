package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pjt_article/Models"
	"pjt_article/Repositories"
	"strconv"
)

func CreateArticle(c *gin.Context) {

	title := c.PostForm("title")
	body := c.PostForm("body")
	userid, _ := strconv.Atoi(c.PostForm("userid"))
	article := Models.Article{Title: title, Body: body, UserID: uint(userid)}
	/*
	var article Models.Article

	c.BindJSON(&article)
	*/
	err := Repositories.CreateArticle(&article)


	if err != nil {

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    &article,
			"message": "created",
		})
	}
}

func GetAllArticles(c *gin.Context){
	var articles []Models.Article
	err := Repositories.GetAllArticles(&articles)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, articles)
	}
}

func GetArticle(c *gin.Context) {
	var article Models.Article
	id := c.Param("id")
	err := Repositories.GetArticle(&article, id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, article)
	}
}

func GetArticleListByUserId(c *gin.Context) {
	var articles []Models.Article
	userid, _ := strconv.Atoi(c.Param("id"))
	err := Repositories.GetArticleListByUserId(&articles, uint(userid))

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, articles)
	}
}

func UpdateArticle(c *gin.Context) {
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

func DeleteArticle(c *gin.Context) {
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
