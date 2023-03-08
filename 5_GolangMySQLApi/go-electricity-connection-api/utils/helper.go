package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/krishnalagad/go-electricity-connection-api/config"
)

var db *gorm.DB

func init() {
	config.ConnectDB()
	db = config.GetDBInstance()
}


