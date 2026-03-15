package main

import (
	"github.com/DidierValdivieso/gin-gorm-rest/config"
	"github.com/DidierValdivieso/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}
