package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Name   string             `json:"name" bson:"name" binding:"required,caontain_alphanum"`
	Price  float64            `json:"price" bson:"price" binding:"required,min=1"`
}
