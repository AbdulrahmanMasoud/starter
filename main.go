package main

import (
	"github.com/AbdulrahmanMasoud/starter/Application"
	"github.com/AbdulrahmanMasoud/starter/Models"
	"github.com/gin-gonic/gin"
)

func main() {

	//هنا اناعرفت متغير وشايل الكلوجر اللي اسمه app

	app := Application.NewApp()        //Start App
	app.DB.AutoMigrate(&Models.User{}) //Make migration
	Application.CloseConnection(&app)  //Close Connection

	//Routing
	app.Gin.GET("/ping", func(c *gin.Context) {
		request := Application.NewRequest(c)
		request.NotAuth()
	})
	app.Gin.Run()
}
