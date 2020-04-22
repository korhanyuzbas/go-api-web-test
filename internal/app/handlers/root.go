package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Sancap"})
	//ctx.JSON(http.StatusOK, gin.H{"message": "hello canımmm"})
}
