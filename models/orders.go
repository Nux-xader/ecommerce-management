package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OrderStatusPending    = "pending"
	OrderStatusProcessing = "processing"
	OrderStatusCompleted  = "completed"
)

type Order struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Products []Product          `json:"products" bson:"products"`
	Status   string             `json:"status" bson:"status,omitempty" binding:"required,order_status"`
}
