package Data

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body string
	UserID uint
	ArticleID uint
}

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Articles []Article
	//Comment []Comment
}

type Article struct {
	gorm.Model
	Title   string
	Body    string
	UserID  uint `json:"user_id"`
	//Comment []Comment
}


