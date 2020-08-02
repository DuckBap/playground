package user

import (
	"GoBoard/database"
	"github.com/gin-gonic/gin"
	"GoBoard/database/models"
	"net/http"
)
var db = database.DB

type User models.User

func CreateUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}