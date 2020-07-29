package Repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pjt_article/Databases"
	"pjt_article/Models"
)


func findArticle(c *gin.Context, db *gorm.DB) Models.Article {
	var article Models.Article
	id := c.Param("id") // id type: string

	db.First(&article, id) // find product with id
	return article
}

func CreateArticle(article *Models.Article) error {
	err := Databases.Db.Create(&article).Error
	if err != nil {
		//Databases.Db.LogMode(true)
		return err
	}
	return nil
}

func GetAllArticles(articles *[]Models.Article) error {
	err := Databases.Db.Find(&articles).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticle(article *Models.Article, id string) error {
	err := Databases.Db.First(&article, id).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticleListByUserId(article *[]Models.Article, userid uint) error {
	err := Databases.Db.Find(&article, "user_id = ?", userid).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}


func UpdateArticle(article *Models.Article, body string) error {
	err := Databases.Db.Model(&article).Update("body", body).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(article *Models.Article) error {
	err := Databases.Db.Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}