package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator_pkg "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationFn func(fl validator_pkg.FieldLevel) bool

var trans ut.Translator
var Validator *validator_pkg.Validate

func init() {
	Validator = validator_pkg.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(Validator, trans)

	AddCustomTag("phoneNumber", "{0} must starts with '62' and have length between 10-14 character", validatePhoneNumber)
}

// Validate accept struct and pointer struct type
func Validate(data interface{}) map[string][]string {
	err := Validator.Struct(data)
	if err == nil {
		return nil
	}
	return translateError(data, err, trans)
}

// errMessage's format must be like this -> "{0} must between '{1}'",
// where {0} is a struct field and {1} is its params,
// it's also can define struct field only -> "{0} must be float".
func AddCustomTag(tag string, errMessage string, validationFn ValidationFn) {
	_ = Validator.RegisterValidation(tag, validator_pkg.Func(validationFn))
	addTranslation(tag, errMessage)
}

// translateError translate result of validation to readable format
func translateError(data interface{}, err error, trans ut.Translator) (result map[string][]string) {
	if err == nil {
		return nil
	}

	result = make(map[string][]string)
	rt := reflect.TypeOf(data)
	if rt.Kind() == reflect.Ptr {
		rt = reflect.ValueOf(data).Elem().Type()
	}
	isDataStruct := rt.Kind() == reflect.Struct

	validatorErrs := err.(validator_pkg.ValidationErrors)
	for _, e := range validatorErrs {
		errorList := make([]string, 0)
		keyName := e.Field()
		translatedErr := errors.New(e.Translate(trans))
		if isDataStruct {
			f, ok := rt.FieldByName(e.Field())
			if jsonVal := strings.TrimSpace(f.Tag.Get("json")); jsonVal != "" && jsonVal != "-" && ok {
				keyName = jsonVal
			}
		}
		if errs, ok := result[keyName]; ok {
			errorList = append(errs, translatedErr.Error())
		} else {
			errorList = append(errorList, translatedErr.Error())
		}
		result[keyName] = errorList
	}
	return result
}

// addTranslation register custom translation of custom validation
func addTranslation(tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator_pkg.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = Validator.RegisterTranslation(tag, trans, registerFn, transFn)
}
