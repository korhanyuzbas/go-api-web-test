package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"sancap/internal/configs"
	"strings"
)

func BindParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}

	validate, err := getValidator(c)
	if err != nil {
		return err
	}

	trans, err := getTranslation(c)
	if err != nil {
		return err
	}
	err = validate.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

func getValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(configs.ValidatorKey)
	if !ok {
		return nil, errors.New("validator is not set")
	}
	validate, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("validator get error")
	}
	return validate, nil
}

func getTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get(configs.TranslatorKey)
	if !ok {
		return nil, errors.New("translator is not set")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("translate get error")
	}
	return translator, nil
}
