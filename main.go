package main

import (
	"GoBoard/database"
	"GoBoard/user"
	"gorm.io/gorm"

	"GoBoard/database/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

const secretKey = "HELLO!MY!NAME!IS!HYEKIM!IT!IS!SECRET!"


func main() {
	r := gin.Default()

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test article",
		Key:         []byte(secretKey),
		IdentityKey: "user",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id": v.ID,
					"username": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Model: gorm.Model{ID:uint(claims["id"].(float64))},
				Username: claims["username"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login Login
			var user models.User
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := login.Username
			password := login.Password
			err := database.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
	})
	r.POST("/sign-up", user.CreateUser)
	r.POST("/login", authMiddleware.LoginHandler)

	//r.POST("/article", article.CreateArticle)
	//r.POST("/article", authMiddleware.MiddlewareFunc(), article.CreateArticle)

	//r.Use(authMiddleware.MiddlewareFunc())

	r.Run(":8080")
}