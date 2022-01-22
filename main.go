package main

import "github.com/gin-gonic/gin"

func main() {

	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	application.Gin.Run()
}
