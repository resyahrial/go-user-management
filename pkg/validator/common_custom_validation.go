package validator

import (
	"regexp"

	validator_pkg "github.com/go-playground/validator/v10"
)

var validatePhoneNumber ValidationFn = func(fl validator_pkg.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	if len(phoneNumber) < 10 || len(phoneNumber) > 14 {
		return false
	}

	phonePattern := regexp.MustCompile(`(?m)(62)[0-9]+$`)
	result := phonePattern.MatchString(phoneNumber)
	return result
}
