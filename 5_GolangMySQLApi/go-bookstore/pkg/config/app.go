package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:krishna24/gobook?charset=utf8&parseTime=True&loc=Local")
	// d, err := sql.Open("mysql", "root:krishna24/gobook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
