package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sancap/internal/configs"
	"sancap/internal/models"
)

func UserMe(ctx *gin.Context) {
	var user models.User
	authid, _ := ctx.Get(configs.JwtKey)
	fmt.Println(authid)
	if err := user.GetByName(authid.(*models.User).Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, user)
}
