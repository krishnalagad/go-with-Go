package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnalagad/gin-gorm-rest/config"
	"github.com/krishnalagad/gin-gorm-rest/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func New() *UserRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{db: db}
}

func UserController(c *gin.Context) {
	c.String(200, "Hello world")
}

// create user
func (userRepo *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(userRepo.db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, user)
}
