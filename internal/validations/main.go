package validations

import (
	"github.com/go-playground/validator/v10"
	"sancap/internal/handlers"
)

func InitValidations(validate *validator.Validate, handler handlers.BaseHandler) {
	userValidations(validate, handler)
}
