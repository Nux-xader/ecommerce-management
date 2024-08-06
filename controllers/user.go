package controllers

import (
	"net/http"
	"time"

	"github.com/Nux-xader/ecommerce-management/config"
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

func ForgotPassword(c *gin.Context) {
	var req models.UserForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	user, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResp("User not found"))
		return
	}

	user.ResetPasswordToken = primitive.NewObjectID().Hex()
	user.ResetPasswordExpires = time.Now().Add(1 * time.Hour)

	if err := repositories.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to update user with reset token"))
		return
	}

	resetLink := config.FRONTEND_RESET_PASSWORD_ROUTE + user.ResetPasswordToken
	htmlContent := "Dear " + user.Username + "<br><br>To reset your password, please click the following link: <a href='" + resetLink + "'>" + resetLink + "</a>"
	go utils.SendEmail(user.Email, "RESET PASSWORD", htmlContent)

	c.JSON(http.StatusOK, utils.SuccessResp())
}

func ResetPassword(c *gin.Context) {
	token := c.Param("token")
	var request models.UserResetPasswowrdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp(err.Error()))
		return
	}

	user, err := repositories.GetUserByResetToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResp("Invalid or expired token"))
		return
	}

	if user.ResetPasswordExpires.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, utils.ErrorResp("Token expired"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to hash password"))
		return
	}

	// change token to expired
	user.ResetPasswordExpires = time.Now().Add(-1 * time.Hour)
	// set new password
	user.Password = string(hashedPassword)

	if err := repositories.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResp("Failed to update password"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResp())
}
