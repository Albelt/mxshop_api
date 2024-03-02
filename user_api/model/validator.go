package model

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	const mobilePattern = "^1[0-9]{10}$"

	mobile, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	ok, _ = regexp.MatchString(mobilePattern, mobile)
	if !ok {
		return false
	}

	return true
}
