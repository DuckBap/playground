package main

import (
	"pjt_article/Databases"
	"pjt_article/routers"
)

func main() {

	//Gorm
	Databases.InitMysql()

	//Gin
	routers.InitRouter()
}
