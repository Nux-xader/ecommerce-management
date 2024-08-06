package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IsContainAlphanum(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return re.MatchString(fl.Field().String())
}
