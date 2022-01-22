package main

import "github.com/gin-gonic/gin"

// ده اللي فيه الحاجات الجلوبل اللي بتشتغل مره واحده

type Application struct {
	Gin *gin.Engine
}

// دي فانكشن بترجع فانكشن
// و الكلوجر ده عشان  اقدر استخدمه في الابلكيشن كله
// وده كلوجر هيبقا جلوبل علي كل الابلكيشن اقدر اكسس منه اللي محتاجه
func app() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		return application
	}
}
