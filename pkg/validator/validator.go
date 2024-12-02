package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/tiancheng92/seminar/pkg/errors/ecode"
	"strings"

	"github.com/Yostardev/gf"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/tiancheng92/seminar/pkg/errors"
)

var translator ut.Translator

func Init() {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("binding validate engine failed")
	}

	zhT := zh.New()
	enT := en.New()
	uni := ut.New(enT, zhT)

	var found bool
	translator, found = uni.GetTranslator("zh")
	if !found {
		panic("translator not found")
	}

	if err := zhTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		panic(gf.StringJoin("register default translations failed: ", err.Error()))
	}

	for i := range customValidators {
		customValidators[i].Register(validate)
	}
}

func HandleValidationErr(err error) error {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		errList := make([]string, 0, len(validationErr))
		for _, v := range validationErr.Translate(translator) {
			errList = append(errList, v)
		}
		return errors.WithCode(ecode.ErrParam, strings.Join(errList, "; "))
	} else {
		return errors.WithCode(ecode.ErrParam, err)
	}
}
