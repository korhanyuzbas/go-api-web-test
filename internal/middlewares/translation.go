package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	trTranslations "github.com/go-playground/validator/v10/translations/tr"
	"log"
	"sancap/internal/configs"
	"sancap/internal/i18n"
	"sancap/internal/validations"
)

func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		enLang := en.New()

		uni := ut.New(enLang, enLang)
		validate := validator.New()

		locale := c.DefaultQuery("locale", "tr")
		trans, _ := uni.GetTranslator(locale)

		if err := trTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
			log.Panicln(err)
		}

		validations.InitValidations(validate)
		i18n.InitTranslations(validate, trans)

		c.Set(configs.TranslatorKey, trans)
		c.Set(configs.ValidatorKey, validate)
		c.Next()
	}
}
