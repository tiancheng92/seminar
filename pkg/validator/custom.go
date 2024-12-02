package validator

import (
	"github.com/Yostardev/gf"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
)

var customValidators = []*validatorInfo{
	{
		Tag:           "date_time_format",
		Kind:          "string",
		Msg:           "{0}日期/时间格式错误(格式: {1})",
		ValidatorFunc: dateTimeFormatValidatorFunc,
	},
}

// defaultTranslateFunc {0} => Field, {1} => Param
func defaultTranslateFunc(translator ut.Translator, fieldError validator.FieldError) string {
	var kindStr string
	kind := fieldError.Kind()
	if kind == reflect.Ptr {
		kind = fieldError.Type().Elem().Kind()
	}

	switch kind {
	case reflect.String:
		kindStr = "string"
	case reflect.Slice, reflect.Map, reflect.Array:
		kindStr = "object"
	default:
		kindStr = "number"
	}
	msg, err := translator.T(gf.StringJoin(fieldError.Tag(), "-", kindStr), fieldError.Field(), fieldError.Param())
	if err != nil {
		panic(gf.StringJoin("register validation failed: ", err.Error()))
	}
	return msg
}

func dateTimeFormatValidatorFunc(fieldLevel validator.FieldLevel) bool {
	date := fieldLevel.Field().String()
	if date == "" {
		return true
	}
	layout := fieldLevel.Param()
	if _, err := time.ParseInLocation(layout, date, time.Local); err != nil {
		return false
	}
	return true
}
