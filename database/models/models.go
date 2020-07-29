package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Body     string
	UserID   uint
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	Body      string
	ArticleID uint `gorm:"not null<-:create"`
	UserID    uint `gorm:"not null<-:create"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
	// Articles []Article
	// Comments []Comment
}