package Repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"pjt_article/Databases"
	"pjt_article/Models"
)


func findUser(c *gin.Context, db *gorm.DB) Models.Article {
	var article Models.Article
	id := c.Param("id") // id type: string

	db.First(&article, id) // find product with id
	return article
}

func CreateUser(user *Models.User) error {
	err := Databases.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(users *[]Models.User) error {
	err := Databases.Db.Find(&users).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(article *Models.Article, id string) error {
	err := Databases.Db.First(&article, id).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(article *Models.Article, body string) error {
	err := Databases.Db.Model(&article).Update("body", body).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(article *Models.Article) error {
	err := Databases.Db.Delete(&article).Error
	if err != nil {
		return err
	}
	return nil
}