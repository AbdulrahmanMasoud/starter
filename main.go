package main

import "github.com/gin-gonic/gin"

func main() {

	//هنا اناعرفت متغير وشايل الكلوجر اللي اسمه app

	app := app()
	application := app() //app ده المتغير مش الفانكشن
	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)
		request.Context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	application.Gin.Run()
}
