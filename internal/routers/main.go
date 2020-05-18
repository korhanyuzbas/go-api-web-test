package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"sancap/internal/handlers"
)

func SetupRouter(db *gorm.DB) (handlers.BaseHandler, *gin.Engine) {
	r := gin.Default()
	handler := handlers.BaseHandler{DB: db}
	SetupAppRouter(r, handler)
	return handler, r
}
