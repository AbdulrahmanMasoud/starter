package Visitors

import (
	"github.com/AbdulrahmanMasoud/starter/Application"
	"github.com/AbdulrahmanMasoud/starter/Models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	request := Application.NewRequest(c).Auth()
	if !request.IsAuth {
		request.NotAuth()
		return
	}
	//Create User
	user := Models.User{
		Username: "Abdulrahman",
		Email:    "abdulrahman@mail.com",
		Password: "masoud",
	}
	request.DB.Create(&user) //to create the user

	request.Created(user) //Return Created Response
}
