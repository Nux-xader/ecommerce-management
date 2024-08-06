package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	jwt.StandardClaims
}

func GenerateJWT(userID primitive.ObjectID) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		fmt.Print(err)
		return nil, nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	// Convert user_id to ObjectID
	if !claims.UserID.IsZero() {
		userID, err := primitive.ObjectIDFromHex(claims.UserID.Hex())
		if err != nil {
			return nil, nil, errors.New("invalid user ID")
		}
		claims.UserID = userID
	}

	return token, claims, nil
}
