package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() {
	data, err := gorm.Open("mysql", "root:krishna24@/gobook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = data
}

func GetDBInstance() *gorm.DB {
	return db
}
