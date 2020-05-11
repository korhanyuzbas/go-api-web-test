package routers

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"sancap/internal/routers/api"
	"sancap/internal/routers/web"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	web.SetupAppRouter(r)
	api.SetupAPIRouter(r)
	return r
}

func CreateDataParams(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}
