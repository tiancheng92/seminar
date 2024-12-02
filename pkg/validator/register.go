package validator

import (
	"github.com/Yostardev/gf"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		return trans.Add(tag, msg, true)
	}
}

type validatorInfo struct {
	Tag           string
	Kind          string
	Msg           string
	TranslateFunc func(ut.Translator, validator.FieldError) string
	ValidatorFunc func(validator.FieldLevel) bool
}

func (v *validatorInfo) Register(validate *validator.Validate) {
	if v.ValidatorFunc != nil {
		if err := validate.RegisterValidation(v.Tag, v.ValidatorFunc); err != nil {
			panic(gf.StringJoin("register validation failed: ", err.Error()))
		}
	}

	if v.TranslateFunc == nil {
		v.TranslateFunc = defaultTranslateFunc
	}

	if err := validate.RegisterTranslation(v.Tag, translator, registerTranslator(gf.StringJoin(v.Tag, "-", v.Kind), v.Msg), v.TranslateFunc); err != nil {
		panic(gf.StringJoin("register validation translation failed: ", err.Error()))
	}
}
