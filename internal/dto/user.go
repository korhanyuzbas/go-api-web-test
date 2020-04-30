package dto

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/utils"
)

type CreateUserInput struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Username  string `json:"username" form:"username" validate:"username-check,required"`
	Password  string `json:"password" form:"password" validate:"required"`
}

func (params *CreateUserInput) ShouldBind(ctx *gin.Context) error {
	return utils.BindParams(ctx, params)
}

type LoginUserInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (params *LoginUserInput) ShouldBind(ctx *gin.Context) error {
	return utils.BindParams(ctx, params)
}

type ChangePasswordInput struct {
	OldPassword  string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword  string `json:"new_password" form:"new_password" validate:"required"`
	NewPassword2 string `json:"new_password2" form:"new_password2" validate:"required,eqfield=NewPassword"`
}

func (params *ChangePasswordInput) ShouldBind(ctx *gin.Context) error {
	return utils.BindParams(ctx, params)
}
