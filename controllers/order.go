package controllers

import (
	"net/http"

	"github.com/Nux-xader/ecommerce-management/models"
	"github.com/Nux-xader/ecommerce-management/repositories"
	"github.com/Nux-xader/ecommerce-management/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrders(c *gin.Context) {
	userID, _ := c.Get("userID")
	products, err := repositories.GetAllOrders(userID.(primitive.ObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	if products == nil {
		products = []models.Order{}
	}
	c.JSON(http.StatusOK, utils.SuccessResp(products))
}

func AddOrder(c *gin.Context) {
	var req models.AddOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	userID, _ := c.Get("userID")
	products, err := repositories.GetAllProducts(userID.(primitive.ObjectID), req.ProductIDs...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed getting products detail"))
		return
	}
	if len(products) == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorResp("products_id must contain at least 1 valid product_id"))
		return
	}

	// copy the product so that it is not affected by future edits or deletes
	order := models.Order{
		UserID:   userID.(primitive.ObjectID),
		Products: products,
		Status:   models.OrderStatusPending,
	}

	err = repositories.AddOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to create order"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResp())
}
