package routes

import (
	"github.com/Nux-xader/ecommerce-management/controllers"
	"github.com/Nux-xader/ecommerce-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	auth := api.Group("/auth")

	user := api.Group("/user")

	products := api.Group("/products")
	products.Use(middleware.JWTMiddleware())

	orders := api.Group("/orders")
	orders.Use(middleware.JWTMiddleware())
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		user.GET("/profile", middleware.JWTMiddleware(), controllers.Profile)
		user.POST("/forgot-password", controllers.ForgotPassword)
		user.PUT("/reset-password/:token", controllers.ResetPassword)

		products.GET("", controllers.GetProducts)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:id", middleware.SlugObjectID("id"), controllers.UpdateProduct)
		products.DELETE("/:id", middleware.SlugObjectID("id"), controllers.DeleteProduct)

		orders.GET("", controllers.GetOrders)
		orders.POST("", controllers.AddOrder)
		orders.PUT("/:id", middleware.SlugObjectID("id"), controllers.SetOrderStatus)
	}
}
