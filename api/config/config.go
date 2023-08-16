package config

import (
	config "erp-system.com/v1/api/middleware"
	"erp-system.com/v1/api/routes"
	"github.com/gin-gonic/gin"
)

func setupAllRoutes(router *gin.Engine, providers ...routes.RouteProvider) {
	for _, provider := range providers {
		provider.SetupRoutes(router)
	}
}

func Setup() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(config.GinLogger())

	authRoutes := &routes.AuthRoutes{}

	setupAllRoutes(r, authRoutes)

	r.Run(":8080")
}
