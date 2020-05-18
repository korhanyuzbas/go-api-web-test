package routers

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers"
	"sancap/internal/middlewares"
)

func setupUserRouter(router *gin.Engine, handler handlers.BaseHandler) *gin.Engine {
	user := router.Group("user")
	user.Use(middlewares.TranslationMiddleware(handler))

	user.GET("login", handler.UserLogin)
	user.GET("register", handler.CreateUser)
	user.POST("register", handler.CreateUser)

	{
		authMiddleware := middlewares.AuthenticationWeb(handler)
		user.POST("login", authMiddleware.LoginHandler)
	}
	{
		user.Use(middlewares.AuthenticationWeb(handler).MiddlewareFunc())
		user.GET("me", handler.UserMe)
		user.GET("change_password", handler.UserChangePassword)
		user.POST("change_password", handler.UserChangePassword)
	}
	return router
}
