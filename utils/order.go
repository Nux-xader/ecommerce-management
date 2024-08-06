package utils

import (
	"github.com/Nux-xader/ecommerce-management/models"
	"github.com/go-playground/validator/v10"
)

var validStatuses = map[string]bool{
	models.OrderStatusPending:    true,
	models.OrderStatusProcessing: true,
	models.OrderStatusCompleted:  true,
}

func OrderStatusValidation(fl validator.FieldLevel) bool {
	return validStatuses[fl.Field().String()]
}
