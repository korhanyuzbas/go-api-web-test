package api

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers/api"
	"sancap/internal/middlewares"
)

func SetupUserAPIRoutes(apiGroup *gin.RouterGroup) {
	userAPIGroup := apiGroup.Group("user")
	userAPIGroup.Use(middlewares.AuthenticationAPI().MiddlewareFunc())
	{
		userAPIGroup.GET("/", api.Index)
		userAPIGroup.GET("/me", api.UserMe)
	}
}
