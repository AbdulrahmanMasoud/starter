package main

import (
	"github.com/AbdulrahmanMasoud/starter/Application"
	"github.com/AbdulrahmanMasoud/starter/Models"
	"github.com/AbdulrahmanMasoud/starter/Routes"
)

func main() {

	//هنا اناعرفت متغير وشايل الكلوجر اللي اسمه app

	app := Application.NewApp()        //Start App
	app.DB.AutoMigrate(&Models.User{}) //Make migration
	Application.CloseConnection(&app)  //Close Connection

	//Routing
	router := Routes.Router{Application: &app} //Start Routing
	router.Routes()                            //Get All Routs
	app.Gin.Run()
}
