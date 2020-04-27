package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
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

		locale := c.DefaultQuery("locale", "en")
		trans, _ := uni.GetTranslator(locale)

		if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
			log.Panicln(err)
		}

		validations.InitValidations(validate)
		i18n.InitTranslations(validate, trans)

		c.Set(configs.TranslatorKey, trans)
		c.Set(configs.ValidatorKey, validate)
		c.Next()
	}
}
