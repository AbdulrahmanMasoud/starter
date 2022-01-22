package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ده عشان  اضيف فيه كل الحاجات اللي محتاجها لما اجي اعمل ريكوس في ال api

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		// هنا في الريكوست بيعمل كونكشن بالداتابيز
		request.DB, request.Connection = connectToDataBase()
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
	req.closeConnection()
	req.Context.JSON(code, body)
}

func connectToDataBase() (*gorm.DB, *sql.DB) {
	dsn := "masoud:password@tcp(127.0.0.1:3306)/gostarter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect to database", err.Error())
	}
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection", err.Error())
	}
	return db, connection

}

// ميثود في الريكوست عشان اقفل الكونكشن بتاع الداتابيز عشان طبعا كل داتا بيز وليها ليمت معين للكونكشن
func (req Request) closeConnection() {
	req.Connection.Close()
}
