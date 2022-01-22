package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeConnection() *gorm.DB {
	dsn := "masoud:password@tcp(127.0.0.1:3306)/gostarter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect to database", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on get database connection", err.Error())
	}
	return connection
}

func connectToDataBase(share ShareResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.DB = makeConnection()
		app.Connection = returnConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = returnConnection(req.DB)
	}
}

// ميثود في الريكوست عشان اقفل الكونكشن بتاع الداتابيز عشان طبعا كل داتا بيز وليها ليمت معين للكونكشن
func closeConnection(share ShareResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.Connection.Close()
	case *Request:
		req := share.(*Request)
		req.Connection.Close()
	}
}
