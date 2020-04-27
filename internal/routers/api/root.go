package api

import (
	"github.com/gin-gonic/gin"
	api2 "sancap/internal/handlers/api"
	"sancap/internal/routers/api/user"
)

func SetupAPIRouter(router *gin.Engine) {
	apiGroup := router.Group("api")
	{
		apiGroup.GET("/", api2.Index)
	}

	user.SetupUserAPIRoutes(apiGroup)

}
