package controllers

import (
	"ginTraining/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	ProductService services.ProductService
}

func (pr *ProductController) AddProduct(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Route",
	})

}

// A product api
// will take the product data and the image of the product
// and the image will be uploaded to aws s3
// And the product data along with the image link
// dockerizing our application
// Makefile in your app
