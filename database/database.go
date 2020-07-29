package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwon/structInfo"
)

var Db *gorm.DB
var Err error

func init() {
	Dsn := "admin:@(localhost)/jwon?charset=utf8&parseTime=True&loc=Local"
	Db, Err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if  Err != nil{
		panic("failed to connect database")
	}
	_ = Db.AutoMigrate(&structInfo.User{}, &structInfo.Article{})
	//userDsn := "admin:@(localhost)/user?charset=utf8&parseTime=True&loc=Local"
	//articleDsn := "admin:@(localhost)/article?charset=utf8&parseTime=True&loc=Local"
	//user.Db, user.Err = gorm.Open(mysql.Open(userDsn), &gorm.Config{})
	//article.Db, article.Err = gorm.Open(mysql.Open(articleDsn), &gorm.Config{})
	//if  user.Err != nil || article.Err != nil{
	//	panic("failed to connect database")
	//}
	//_ = user.Db.AutoMigrate(&structInfo.User{}, &structInfo.Article{})
	//_ = article.Db.AutoMigrate(&structInfo.Article{})
}
