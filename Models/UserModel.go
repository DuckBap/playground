package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	Articles []Article
}