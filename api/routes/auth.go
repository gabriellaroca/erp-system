package routes

import (
	"erp-system.com/v1/adapter/mongodb"
	"erp-system.com/v1/adapter/mongodb/config"
	"erp-system.com/v1/api"
	"erp-system.com/v1/domain/usecase"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct{}

func (u *AuthRoutes) SetupRoutes(router *gin.Engine) {
	personRepo := mongodb.NewPersonRepository(config.MongoClient.Database("erp-system").Collection("Person"))
	personUsecase := usecase.NewPersonUsecase(personRepo)
	authHandler := api.NewAuthHandler(personUsecase)

	userGroup := router.Group("/auth")
	{
		userGroup.POST("/login", authHandler.Login)
		userGroup.GET("/access", authHandler.GetAcess)
	}
}

func SetRoute() {
}
