package models

import (
	"github.com/krishnalagad/gin-gorm-rest/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	City        string `json:"city"`
	Email       string `json:"email"`
	Designation string `json:"education"`
}

func init() {
	db = config.InitDb()
	db.AutoMigrate(&User{})
}

func CreateUser(db *gorm.DB, user *User) *User {
	db.Create(user)
	return user
}
