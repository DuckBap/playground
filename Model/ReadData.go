package Model

import (
	"encoding/json"
	"example.com/test/Data"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"net/http"
	"strconv"
)
//
//func ReadCommentData (c *gin.Context) {
//	var comments []Comment
//
//	SId := c.Query("id")
//	fmt.Println(SId)
//	id,_ := strconv.Atoi(SId)
//	comId := uint(id)
//	ReadDataBase(&comments,"user_id", comId)
//	datas,_ := json.MarshalIndent(comments, "", "	")
//	c.String(http.StatusOK, "%s\n", datas)
//}

func ReadArticleData (c *gin.Context) {
	var articles []Data.Article

	SId := c.Query("id")
	fmt.Println(SId)
	id,_ := strconv.Atoi(SId)
	comId := uint(id)
	ReadDataBase(&articles,"id", comId)
	datas,_ := json.MarshalIndent(articles, "", "	")
	c.String(http.StatusOK, "%s\n", datas)
}

func ReadUserData (c *gin.Context) {
	var users []Data.User

	SId := c.Query("id")
	fmt.Println(SId)
	id,_ := strconv.Atoi(SId)
	comId := uint(id)
	ReadDataBase(&users,"id", comId)
	datas,_ := json.MarshalIndent(users, "", "	")
	c.String(http.StatusOK, "%s\n", datas)
}

func ReadDataBase (products interface{}, condition string, id uint) {
	query := condition + " = ?"
	dsn := "dokang:1234@tcp(127.0.0.1:3306)/exdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	mysqlDB, err := db.DB()
	if err != nil {
		panic("error")
	}
	defer mysqlDB.Close()
	db.Where(query,id).Find(products)

}