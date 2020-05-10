package web

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/configs"
	"sancap/internal/handlers/web"
)

func SetupAppRouter(router *gin.Engine, withTemplates bool) {
	if withTemplates {
		router.LoadHTMLGlob(configs.AppConfig.TemplateDir)
	}
	router.GET("/", web.Index)

	setupUserRouter(router)
}
