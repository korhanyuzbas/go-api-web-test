package web

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers/web"
	"sancap/internal/middlewares"
)

func setupUserRouter(router *gin.Engine) *gin.Engine {
	user := router.Group("user")
	user.Use(middlewares.TranslationMiddleware())

	user.GET("login", web.UserLogin)
	user.POST("login", web.UserLogin)
	user.GET("register", web.CreateUser)
	user.POST("register", web.CreateUser)

	{
		user.Use(middlewares.AuthenticationWeb().MiddlewareFunc())
		user.GET("me", web.UserMe)
		user.GET("change_password", web.UserChangePassword)
		user.POST("change_password", web.UserChangePassword)
	}
	return router
}
