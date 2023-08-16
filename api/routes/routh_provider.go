package routes

import "github.com/gin-gonic/gin"

type RouteProvider interface {
	SetupRoutes(router *gin.Engine)
}
