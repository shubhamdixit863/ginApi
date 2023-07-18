package controllers

import (
	"ginTraining/internal/dtos"
	"ginTraining/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {

	// Service as a depeendency here

	AuthService *services.AuthService
}

func (at *AuthController) Signup(c *gin.Context) {
	var signupRequest dtos.SignupRequest
	err := c.ShouldBindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})

		return
	}

	/// We can use service methods here
	token, err := at.AuthService.SignupUser(signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signnup Success",
		"token":   token,
	})

}

func (at *AuthController) Login(c *gin.Context) {

	var loginRequest dtos.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})

		return
	}

	/// We can use service methods here
	token, err := at.AuthService.LoginUser(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"token":   token,
	})

}
