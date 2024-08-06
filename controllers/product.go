package controllers

import (
	"net/http"

	"github.com/Nux-xader/ecommerce-management/models"
	"github.com/Nux-xader/ecommerce-management/repositories"
	"github.com/Nux-xader/ecommerce-management/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c *gin.Context) {
	userID, _ := c.Get("userID")
	products, err := repositories.GetAllProducts(userID.(primitive.ObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	if products == nil {
		products = []models.Product{}
	}
	c.JSON(http.StatusOK, utils.SuccessResp(products))
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}
	userID, _ := c.Get("userID")
	product.UserID = userID.(primitive.ObjectID)

	err := repositories.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResp())
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	productID, _ := c.Get("id")
	userID, _ := c.Get("userID")
	isSuccess, err := repositories.UpdateProduct(
		productID.(primitive.ObjectID),
		userID.(primitive.ObjectID),
		&product,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResp())
}

func DeleteProduct(c *gin.Context) {
	productID, _ := c.Get("id")
	userID, _ := c.Get("userID")
	isSuccess, err := repositories.DeleteProduct(
		productID.(primitive.ObjectID),
		userID.(primitive.ObjectID),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	if !isSuccess {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResp())
}
