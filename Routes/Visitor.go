package Routes

import (
	"github.com/AbdulrahmanMasoud/starter/Controllers/Visitors"
)

func (router Router) visitorsRoute() {
	router.Gin.GET("/user/create", Visitors.CreateUser)
}
