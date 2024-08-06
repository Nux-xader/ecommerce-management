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

	products := api.Group("/products")
	products.Use(middleware.JWTMiddleware())
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/forgot_password", controllers.ForgotPassword)
		auth.POST("/reset_password/:token", controllers.ResetPassword)

		profile.GET("", controllers.Profile)

		products.GET("", controllers.GetProducts)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:id", middleware.SlugObjectID("id"), controllers.UpdateProduct)
		products.DELETE("/:id", middleware.SlugObjectID("id"), controllers.DeleteProduct)
	}
}
