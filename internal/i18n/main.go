package i18n

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func InitTranslations(validate *validator.Validate, trans ut.Translator) {
	userTranslations(validate, trans)
}
