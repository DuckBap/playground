package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"GoBoard/database/models"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "hyekim:1234@(localhost)/board_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
}
