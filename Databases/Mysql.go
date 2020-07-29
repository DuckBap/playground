package Databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pjt_article/Models"
)

var Db *gorm.DB
var err error

func InitMysql() {
	dsn := "eun-park:1234@/study_db?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		migrateMysql()
	}
}

func migrateMysql() {
	Db.AutoMigrate(&Models.User{}, &Models.Article{})
}
