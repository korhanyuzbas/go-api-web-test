package dto

import (
	"github.com/gin-gonic/gin"
	"sancap/internal/configs"
)

type CreateUserInput struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Username  string `json:"username" form:"username" binding:"required" validate:"username-check,required"`
	Password  string `json:"password" form:"password" json:"password"`
}

func (params *CreateUserInput) ShouldBind(c *gin.Context) error {
	return configs.BindParams(c, params)
}
