package api

import (
	"github.com/gin-gonic/gin"
	api2 "sancap/internal/handlers/api"
)

func SetupAPIRouter(router *gin.Engine) {
	apiGroup := router.Group("api")
	{
		apiGroup.GET("/", api2.Index)
	}

	SetupUserAPIRoutes(apiGroup)

}
