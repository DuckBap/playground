package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model //id, create, update
	Title string
	Body string
	UserID uint
}
