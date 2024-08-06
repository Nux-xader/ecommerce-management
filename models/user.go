package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username             string             `bson:"username" json:"username" binding:"required,alphanum"`
	Email                string             `bson:"email" json:"email" binding:"required,email_validation"`
	Password             string             `bson:"password" json:"password" binding:"required,min=8"`
	ResetPasswordToken   string             `bson:"reset_password_token,omitempty" json:"reset_password_token,omitempty"`
	ResetPasswordExpires time.Time          `bson:"reset_password_expires,omitempty" json:"reset_password_expires,omitempty"`
}

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserTokenResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email_validation"`
}

type UserResetPasswowrdRequest struct {
	Password string `json:"password" binding:"required,min=8"`
}
