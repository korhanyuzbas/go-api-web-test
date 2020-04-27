package user

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers/api"
)

func SetupUserAPIRoutes(apiGroup *gin.RouterGroup) {
	userAPIGroup := apiGroup.Group("user")
	{
		userAPIGroup.GET("/", api.Index)
	}
}
