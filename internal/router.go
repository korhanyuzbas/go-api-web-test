package internal

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/routers/api"
	"sancap/internal/routers/web"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	web.SetupAppRouter(r)
	api.SetupAPIRouter(r)
	return r
}
