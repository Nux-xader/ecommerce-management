package controllers

import (
	"net/http"

	"github.com/Nux-xader/ecommerce-management/models"
	"github.com/Nux-xader/ecommerce-management/repositories"
	"github.com/Nux-xader/ecommerce-management/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	// Check if username is already taken
	isUsernameTaken, err := repositories.IsUsernameTaken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to check username"))
		return
	}
	if isUsernameTaken {
		c.JSON(http.StatusBadRequest, utils.ErrorResp("Username is already taken"))
		return
	}

	// Check if email is already taken
	isEmailTaken, err := repositories.IsEmailTaken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to check email"))
		return
	}
	if isEmailTaken {
		c.JSON(http.StatusBadRequest, utils.ErrorResp("Email is already taken"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	user.Password = string(hashedPassword)

	user.ID = primitive.NewObjectID()
	if err := repositories.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResp())
}

func Login(c *gin.Context) {
	var credentials models.LoginCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	user, err := repositories.GetUserByUsername(credentials.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorResp("Invalid username or password"))
		return
	}

	if !utils.IsCorrectPassword(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, utils.ErrorResp("Invalid username or password"))
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResp(models.UserTokenResponse{Token: token}))
}

func Profile(c *gin.Context) {
	userID, _ := c.Get("userID")
	profile, err := repositories.GetUserByID(userID.(primitive.ObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResp(models.UserProfileResponse{
		ID:       profile.ID.Hex(),
		Username: profile.Username,
		Email:    profile.Email,
	}))
}
