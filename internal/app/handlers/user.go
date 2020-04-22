package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sancap/internal/app/models"
)

func CreateUser(ctx *gin.Context) {
	var user models.UserModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err := models.CreateUser(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, user)
		}
	}
}

func UserMe(ctx *gin.Context) {
	var user models.UserModel
	authid, _ := ctx.Get("id")
	if err := models.GetUserByName(&user, authid.(*models.UserModel).Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, user)
}
