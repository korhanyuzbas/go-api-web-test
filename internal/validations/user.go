package validations

import (
	"github.com/go-playground/validator/v10"
	"log"
	"sancap/internal/dto"
	"sancap/internal/handlers"
	"sancap/internal/models"
)

func userValidations(validate *validator.Validate, handler handlers.BaseHandler) {
	if err := validate.RegisterValidation("username-check", func(fl validator.FieldLevel) bool {
		userInput := fl.Parent().Interface().(*dto.CreateUserInput)
		user := &models.User{
			FirstName: userInput.FirstName,
			LastName:  userInput.LastName,
			Username:  userInput.Username,
			Password:  []byte(userInput.Password),
		}
		name := fl.Field().String()
		if user.IsUsernameAvailable(handler.DB, name) {
			return true
		}
		return false
	}); err != nil {
		log.Panicln(err)
	}
}
