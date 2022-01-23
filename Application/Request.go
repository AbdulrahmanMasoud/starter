package Application

import (
	"database/sql"
	"github.com/AbdulrahmanMasoud/starter/Models"
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
	User       *Models.User
	IsAuth     bool
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

func NewRequest(c *gin.Context) *Request {
	request := req()  //جبت الكلوجر بتاعي
	req := request(c) // وبعدين  استخدمته واديتله اللي هو محتاجه
	return &req
}

// الميثود دي عشان  اوحد الريسبونس اللي راجع

func (req Request) Response(code int, body interface{}) {
	//هنا اما يرجع الريسبونس يقف الكونكشن بتاع الداتابيز الاول
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

// الميثود دي عشان تتشك هل اليوزر اللي داخل ده مسجل دخول ولا لا

func (req *Request) Auth() *Request {
	req.IsAuth = false
	authHeader := req.Context.GetHeader("Authorization")
	if authHeader != "" {
		req.DB.Where("token = ?", authHeader).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
		}
	}
	return req
}
