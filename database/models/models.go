package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `form:"title" json:"title" gorm:"not null"`
	Body     string `form:"body" json:"body"`
	UserID   uint	`form:"user_id" json:"user_id"`
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
	Username string `form:"username" json:"username" gorm:"unique"`
	Password string	`form:"password" json:"password"`
	Nickname string	`form:"nickname" json:"nickname"`
	Articles []Article
	// Comments []Comment
}