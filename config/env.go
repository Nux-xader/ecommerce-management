package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGO_URI     string
	MONGO_DB_NAME string

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

	SECRET_KEY = os.Getenv("SECRET_KEY")
	SERVER_ADDRESS = os.Getenv("SERVER_ADDRESS")

	FRONTEND_RESET_PASSWORD_ROUTE = os.Getenv("FRONTEND_RESET_PASSWORD_ROUTE")
}
