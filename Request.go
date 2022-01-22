package main

import "github.com/gin-gonic/gin"

// ده عشان  اضيف فيه كل الحاجات اللي محتاجها لما اجي اعمل ريكوس في ال api

type Request struct {
	Context *gin.Context
}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		return request
	}
}
