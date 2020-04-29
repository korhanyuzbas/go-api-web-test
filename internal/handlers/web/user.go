package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sancap/internal/configs"
	"sancap/internal/dto"
	"sancap/internal/models"
)

func CreateUser(ctx *gin.Context) {
	// TODO: is it good practice?
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "user.register.html", gin.H{"user": nil, "error": nil})
		return
	}
	addInput := &dto.CreateUserInput{}
	if err := addInput.ShouldBind(ctx); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.register.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	user := &models.User{
		FirstName: addInput.FirstName,
		LastName:  addInput.LastName,
		Username:  addInput.Username,
		Password:  addInput.Password,
	}
	if err := user.Create(); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.register.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	ctx.HTML(http.StatusCreated, "user.register.html", gin.H{"user": user, "error": nil})

}

func UserMe(ctx *gin.Context) {
	var user models.User
	authid, _ := ctx.Get(configs.JwtKey)
	if err := user.GetByName(authid.(*models.User).Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, user)
}
