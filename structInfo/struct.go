package structInfo

import (
	"gorm.io/gorm"
)

type User		struct {
	gorm.Model
	Username 	string
	Password	string
	Nickname	string
	Articles 	[]Article
	//Comments	[]Comment
}

type Comment	struct {
	gorm.Model
	Body  		string
}

type Article	struct {
	gorm.Model
	Title		string
	Body 		string
	UserID		uint
	//Comments	[]Comment
}
