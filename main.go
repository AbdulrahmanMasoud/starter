package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//هنا اناعرفت متغير وشايل الكلوجر اللي اسمه app

	app := app()
	application := app() //app ده المتغير مش الفانكشن
	application.Gin.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)
		request.closeConnection() //دي عشان  اقفل الكونكشن بتاع الداتابيز بعد الركوست
		request.Context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	application.Gin.Run()
}
