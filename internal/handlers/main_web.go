package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h BaseHandler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Sancap"})
}
