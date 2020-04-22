package routers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sancap/internal/app/handlers"
	"sancap/internal/configs"
)

func SetupRouter(m *jwt.GinJWTMiddleware) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())
	router.LoadHTMLGlob(configs.AppConfig.TemplateDir)

	router.GET("/", handlers.Index)

	return router
}
