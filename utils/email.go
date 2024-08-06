package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func EmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	re := regexp.MustCompile(EmailRegex)
	return re.MatchString(email)
}
