package article

import (
	"GoBoard/database"
	"GoBoard/database/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Article models.Article
type Comment models.Comment

var db = database.DB

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	userId, _ := strconv.Atoi(c.PostForm("userId"))

	article := Article{Title: title, Body: body, UserID: uint(userId)}
	db.Create(&article)
	c.JSON(200, article)
}

func getArticle(c *gin.Context) {
	var article Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, "")
		return
	}
	err = db.First(&article, id).Error
	if err != nil {
		c.JSON(404, "")
		return
	}
	db.Model(&article).Association("Comments").Find(&article.Comments)
	c.JSON(200, article)
}

func listArticles(c *gin.Context) {
	var articles []Article
	db.Find(&articles)
	c.JSON(200, articles)
}

func updateArticle(c *gin.Context) {
	var article Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, "")
	}
	db.Find(&article, id)
	title, hasKey := c.GetPostForm("title")
	if hasKey {
		article.Title = title
	}
	body, hasKey := c.GetPostForm("body")
	if hasKey {
		article.Body = body
	}
	db.Save(&article)
	c.JSON(200, article)
}

func deleteArticle(c *gin.Context) {
	var article Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, "")
	}
	db.Find(&article, id)
	db.Delete(&article)
	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})
}

func createComment(c *gin.Context) {
	var comment Comment
	var article Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, "")
		return
	}
	err = db.Find(&article, id).Error
	if err != nil {
		c.JSON(404, "")
		return
	}
	comment.Body = c.PostForm("body")
	userIdStr, hasKey := c.GetPostForm("userId")
	if !hasKey {
		c.JSON(400, "")
		return
	}
	userId, _ := strconv.Atoi(userIdStr)
	comment.UserID = uint(userId)
	articleIdStr, hasKey := c.GetPostForm("articleId")
	if !hasKey {
		c.JSON(400, "")
		return
	}
	articleId, _ := strconv.Atoi(articleIdStr)
	comment.ArticleID = uint(articleId)
	db.Create(&comment)
	c.JSON(200, comment)
}
