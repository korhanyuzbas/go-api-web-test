package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sancap/internal/handlers"
)

func SetupRouter(m *jwt.GinJWTMiddleware) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", handlers.Index)

	return router
}
