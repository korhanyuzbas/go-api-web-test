package web

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/configs"
	"sancap/internal/handlers/web"
	"sancap/internal/middlewares"
)

func SetupAppRouter(router *gin.Engine) {
	router.LoadHTMLGlob(configs.AppConfig.TemplateDir)

	router.GET("/", web.Index)

	userGroup := router.Group("user")
	userGroup.Use(middlewares.TranslationMiddleware())
	userGroup.GET("register", web.AddUser)
	userGroup.POST("register", web.AddUser)
}
