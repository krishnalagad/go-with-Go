package models

import (
	"github.com/jinzhu/gorm"
	"github.com/krishnalagad/go-electricity-connection-api/config"
)

var db *gorm.DB

type Conenction struct {
	gorm.Model
	ConnectionOwnerName string `gorm:""json:"name"`
	ConnectionAddress   string `json:"address"`
	ConnectionType      string `json:"type"`
	ConnectionCharges   string `json:"charges"`
}

func init() {
	config.ConnectDB()
	db = config.GetDBInstance()
	db.AutoMigrate(&Conenction{})
}
