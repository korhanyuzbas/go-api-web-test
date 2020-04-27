package i18n

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
)

func userTranslations(validate *validator.Validate, trans ut.Translator) {
	if err := validate.RegisterTranslation("username-check", trans, func(ut ut.Translator) error {
		return ut.Add("username-check", "Username is already exists", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("username-check", fe.Field())
		return t
	}); err != nil {
		log.Panicln(err)
	}
}
