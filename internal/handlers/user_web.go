package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sancap/internal/configs"
	"sancap/internal/dto"
	"sancap/internal/models"
)

func (h BaseHandler) CreateUser(ctx *gin.Context) {
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
	fmt.Println(h.DB)
	user.Create(h.DB)

	ctx.HTML(http.StatusCreated, "user.register.html", gin.H{"user": user, "error": nil})
}

func (h BaseHandler) UserLogin(ctx *gin.Context) {
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
	if err := user.GetCredentials(h.DB, loginInput.Password); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.login.html", gin.H{"error": err.Error(), "user": nil})
		return
	}
	//middlewares.AuthenticationWeb(h).LoginHandler(ctx) // set cookie for web user with custom login handler

	ctx.HTML(http.StatusOK, "user.login.html", gin.H{"error": nil, "user": user})
}

func (h BaseHandler) UserMe(ctx *gin.Context) {
	authid, _ := ctx.Get(configs.JwtKey)
	user := authid.(*models.User)
	ctx.HTML(http.StatusOK, "user.me.html", gin.H{"user": user})
}

func (h BaseHandler) UserChangePassword(ctx *gin.Context) {
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
	if err := user.GetCredentials(h.DB, changePasswordInput.OldPassword); err != nil {
		ctx.HTML(http.StatusBadRequest, "user.change_password.html", gin.H{"user": user, "error": err.Error()})
		return
	}
	user.ChangePassword(h.DB, changePasswordInput.NewPassword)

	ctx.HTML(http.StatusOK, "user.change_password.html", gin.H{"user": user, "error": nil})
}
