package main

import (
	"github.com/Nux-xader/ecommerce-management/config"
	"github.com/Nux-xader/ecommerce-management/repositories"
	"github.com/Nux-xader/ecommerce-management/routes"
	"github.com/Nux-xader/ecommerce-management/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	// Database initialization
	config.LoadConfig()
	config.InitDB()
	repositories.InitUserRepository(config.DB)
	repositories.InitProductRepository(config.DB)
	repositories.InitOrderRepository(config.DB)

	router := gin.Default()

	// Config CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // allow all
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	// set custom binding validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("email_validation", utils.EmailValidation)
		v.RegisterValidation("caontain_alphanum", utils.IsContainAlphanum)
		v.RegisterValidation("order_status", utils.OrderStatusValidation)
	}

	routes.SetupRoutes(router)
	router.Run(config.SERVER_ADDRESS)
}
