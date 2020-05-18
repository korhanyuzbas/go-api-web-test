package routers

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/configs"
	"sancap/internal/handlers"
)

func SetupAppRouter(router *gin.Engine, handler handlers.BaseHandler) {
	router.LoadHTMLGlob(configs.AppConfig.TemplateDir)

	setupUserRouter(router, handler)
	router.GET("/", handler.Index)
}
