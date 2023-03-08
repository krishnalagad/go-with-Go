package models

import "github.com/jinzhu/gorm"

type Conenction struct {
	gorm.Model
	ConnectionOwnerName string `gorm:""json:"name"`
	ConnectionAddress   string `json:"address"`
	ConnectionType      string `json:"type"`
	ConnectionCharges   string `json:"charges"`
}
