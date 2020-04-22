package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sancap/internal/app/handlers"
)

func setupUserRouter(router *gin.Engine, m *jwt.GinJWTMiddleware) *gin.Engine {
	router.POST("register", handlers.CreateUser)
	user := router.Group("user")
	user.Use(m.MiddlewareFunc())
	{
		user.GET("me", handlers.UserMe)
	}
	return router
}
