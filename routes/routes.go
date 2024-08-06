package routes

import (
	"github.com/Nux-xader/ecommerce-management/controllers"
	"github.com/Nux-xader/ecommerce-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	auth := api.Group("/auth")

	profile := api.Group("/profile")
	profile.Use(middleware.JWTMiddleware())

	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		profile.GET("", controllers.Profile)
	}
}
