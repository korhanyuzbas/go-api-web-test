package validations

import "github.com/go-playground/validator/v10"

func InitValidations(validate *validator.Validate) {
	userValidations(validate)
}
