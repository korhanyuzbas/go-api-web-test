package web

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers/web"
)

func setupUserRouter(router *gin.Engine, m *jwt.GinJWTMiddleware) *gin.Engine {
	user := router.Group("user")
	user.Use(m.MiddlewareFunc())
	{
		user.GET("me", web.UserMe)
	}
	return router
}
