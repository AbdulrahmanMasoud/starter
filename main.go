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
	app.Gin.GET("/user/create", func(c *gin.Context) {
		request := Application.NewRequest(c)
		//Create User
		user := Models.User{
			Username: "Abdulrahman",
			Email:    "abdulrahman@mail.com",
			Password: "masoud",
		}
		request.DB.Create(&user) //to create the user

		request.Created(user) //Return Created Response

	})
	app.Gin.Run()
}
