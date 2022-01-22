package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShareResources interface {
	Share()
}

// ده عشان  اضيف فيه كل الحاجات اللي محتاجها لما اجي اعمل ريكوس في ال api

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

func (req *Request) Share() {}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		connectToDataBase(&request)
		return request
	}
}

func newRequest(c *gin.Context) Request {
	request := req()  //جبت الكلوجر بتاعي
	req := request(c) // وبعدين  استخدمته واديتله اللي هو محتاجه
	return req
}

// الميثود دي عشان  اوحد الريسبونس اللي راجع

func (req Request) Response(code int, body interface{}) {
	//هنا اما يرجع الريسبونس يقف الكونكشن بتاع الداتابيز الاول
	closeConnection(&req)
	req.Context.JSON(code, body)
}
