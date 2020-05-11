package web

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/configs"
	"sancap/internal/handlers/web"
)

func SetupAppRouter(router *gin.Engine) {
	router.LoadHTMLGlob(configs.AppConfig.TemplateDir)

	router.GET("/", web.Index)

	setupUserRouter(router)
}
