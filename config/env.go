package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGO_URI     string
	MONGO_DB_NAME string

	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USERNAME string
	SMTP_PASSWORD string
	EMAIL_FROM    string

	SECRET_KEY     string
	SERVER_ADDRESS string

	FRONTEND_RESET_PASSWORD_ROUTE string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	MONGO_DB_NAME = os.Getenv("MONGO_DB_NAME")

	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP PORT")
	}
	SMTP_USERNAME = os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	EMAIL_FROM = os.Getenv("EMAIL_FROM")

	SECRET_KEY = os.Getenv("SECRET_KEY")
	SERVER_ADDRESS = os.Getenv("SERVER_ADDRESS")

	FRONTEND_RESET_PASSWORD_ROUTE = os.Getenv("FRONTEND_RESET_PASSWORD_ROUTE")
}
