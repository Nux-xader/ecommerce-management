package utils

import (
	"regexp"

	"github.com/Nux-xader/ecommerce-management/config"
	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
)

const EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func EmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	re := regexp.MustCompile(EmailRegex)
	return re.MatchString(email)
}

func SendEmail(to, resetLink, subject, htmlContent string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.EMAIL_FROM)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "ECOMMERCE MANAGEMENT - "+subject)
	m.SetBody("text/html", htmlContent)

	d := gomail.NewDialer(
		config.SMTP_HOST,
		config.SMTP_PORT,
		config.SMTP_USERNAME,
		config.SMTP_PASSWORD,
	)

	d.DialAndSend(m)
}
