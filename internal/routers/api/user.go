package api

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/middlewares"
)

func SetupUserAPIRoutes(apiGroup *gin.RouterGroup) {
	userAPIGroup := apiGroup.Group("user")
	userAPIGroup.Use(middlewares.AuthenticationAPI().MiddlewareFunc())
}
