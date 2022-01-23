package Application

import "github.com/gin-gonic/gin"

func (req Request) Ok(body interface{}) {
	req.Response(200, body)
}

func (req Request) Created(body interface{}) {
	req.Response(201, body)
}

func (req Request) NotAuth() {
	req.Response(401, gin.H{
		"message": "You are not Authorized",
	})
}
