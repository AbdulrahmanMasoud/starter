package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50)"`
	Email    string `json:"email" gorm:"type:varchar(50)"`
	Password string `json:"password" gorm:"type:varchar(50)"`
}
