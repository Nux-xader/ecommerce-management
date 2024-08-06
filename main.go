package main

import (
	"os"

	"github.com/Nux-xader/ecommerce-management/config"
	"github.com/Nux-xader/ecommerce-management/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	// Database initialization
	config.InitDB()

	router := gin.Default()

	// Config CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // allow all
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// set custom binding validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("email_validation", utils.EmailValidation)
	}

	router.Run(os.Getenv("SERVER_ADDRESS"))
}
