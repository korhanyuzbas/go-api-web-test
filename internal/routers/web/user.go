package web

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers/web"
	"sancap/internal/middlewares"
)

func setupUserRouter(router *gin.Engine) *gin.Engine {
	user := router.Group("user")
	{
		user.Use(middlewares.TranslationMiddleware())
		user.GET("register", web.CreateUser)
		user.POST("register", web.CreateUser)
	}
	return router
}
