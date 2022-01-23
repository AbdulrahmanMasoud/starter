package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ده اللي فيه الحاجات الجلوبل اللي بتشتغل مره واحده

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share() {}

// دي فانكشن بترجع فانكشن
// و الكلوجر ده عشان  اقدر استخدمه في الابلكيشن كله
// وده كلوجر هيبقا جلوبل علي كل الابلكيشن اقدر اكسس منه اللي محتاجه

func App() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		return application
	}
}

func NewApp() Application {
	app := App()         //جبت الكلوجر بتاعي
	application := app() // وبعدين  استخدمته واديتله اللي هو محتاجه
	return application
}
