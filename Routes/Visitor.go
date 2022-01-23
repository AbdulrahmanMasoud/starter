package Routes

import (
	"github.com/AbdulrahmanMasoud/starter/Application"
	"github.com/AbdulrahmanMasoud/starter/Models"
	"github.com/gin-gonic/gin"
)

func (router Router) VisitorsRoute() {
	router.Gin.GET("/user/create", func(c *gin.Context) {
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
}
