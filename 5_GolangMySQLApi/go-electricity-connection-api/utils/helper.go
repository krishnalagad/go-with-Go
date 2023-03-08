package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/krishnalagad/go-electricity-connection-api/models"
)

var db *gorm.DB

func (c *models.Conenction) createConnection() interface  {
	db.NewRecord(c)
	db.Create(&c)
	return c
}




