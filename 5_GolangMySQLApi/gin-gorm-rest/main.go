package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krishnalagad/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()

	routes.UserRoute(router)

	router.Run(":4000")
}
