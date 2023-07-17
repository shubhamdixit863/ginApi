package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
}

func (pr *ProductController) AddProduct(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Route",
	})

}
