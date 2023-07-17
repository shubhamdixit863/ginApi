package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (at *AuthController) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Signnup Route",
	})

}

func (at *AuthController) Login(c *gin.Context) {

}