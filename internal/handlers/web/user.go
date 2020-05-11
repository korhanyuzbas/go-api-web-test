package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sancap/internal/configs"
	"sancap/internal/dto"
	"sancap/internal/middlewares"
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
		Password:  []byte(addInput.Password),
	}
	if err := user.Create(); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.register.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	ctx.HTML(http.StatusCreated, "user.register.html", gin.H{"user": user, "error": nil})
}

func UserLogin(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "user.login.html", gin.H{"user": nil, "error": nil})
		return
	}
	loginInput := &dto.LoginUserInput{}
	if err := loginInput.ShouldBind(ctx); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.login.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	user := &models.User{
		Username: loginInput.Username,
		Password: []byte(loginInput.Password),
	}
	if err := user.GetCredentials(loginInput.Password); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.login.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	middlewares.AuthenticationWeb().LoginHandler(ctx) // set cookie for web user with custom login handler
	ctx.HTML(http.StatusOK, "user.login.html", gin.H{"error": nil, "user": user})
}

func UserMe(ctx *gin.Context) {
	authid, _ := ctx.Get(configs.JwtKey)
	user := authid.(*models.User)
	ctx.HTML(http.StatusOK, "user.me.html", gin.H{"user": user})
}

func UserChangePassword(ctx *gin.Context) {
	authid, _ := ctx.Get(configs.JwtKey)
	user := authid.(*models.User)
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "user.change_password.html", gin.H{"user": user, "error": nil})
		return
	}
	changePasswordInput := &dto.ChangePasswordInput{}
	if err := changePasswordInput.ShouldBind(ctx); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.change_password.html", gin.H{"error": err.Error(), "user": user})
		return
	}
	if err := user.GetCredentials(changePasswordInput.OldPassword); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.change_password.html", gin.H{"user": user, "error": err.Error()})
		return
	}
	if err := user.ChangePassword(changePasswordInput.NewPassword); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.change_password.html", gin.H{"user": user, "error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "user.change_password.html", gin.H{"user": user, "error": nil})
}
