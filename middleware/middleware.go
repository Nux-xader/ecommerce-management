package middleware

import (
	"net/http"
	"strings"

	"github.com/Nux-xader/ecommerce-management/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.ErrorResp("Authorization header is required"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, utils.ErrorResp("Invalid authorization header format"))
			c.Abort()
			return
		}

		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, utils.ErrorResp("Invalid token"))
			c.Abort()
			return
		}

		userID := claims.UserID
		if userID == primitive.NilObjectID {
			c.JSON(http.StatusUnauthorized, utils.ErrorResp("Invalid user ID in token"))
			c.Abort()
			return
		}

		// Set user ID in context as ObjectID
		c.Set("userID", userID)
		c.Next()
	}
}

func SlugObjectID(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param(param)
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusNotFound, utils.ErrorResp("Invalid Object ID"))
			c.Abort()
			return
		}
		c.Set(param, objectID)
		c.Next()
	}
}
