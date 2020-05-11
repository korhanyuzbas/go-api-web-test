package api

import (
	"github.com/gin-gonic/gin"
)

func SetupAPIRouter(router *gin.Engine) {
	apiGroup := router.Group("api")
	{
		SetupUserAPIRoutes(apiGroup)
	}

}
