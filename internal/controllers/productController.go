package controllers

import (
	"ginTraining/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type ProductController struct {
	ProductService services.ProductService
}

func (pr *ProductController) AddProduct(c *gin.Context) {
	// form data we will be dealing with
	// We will upload the file in local directory
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	// We will upload it in the local directory
	uploadPath := "../../uploads"

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			return
		}
	}

	err = c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}
	//

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"error":   nil,
	})

}

// A product api
// will take the product data and the image of the product
// and the image will be uploaded to aws s3
// And the product data along with the image link
// dockerizing our application
// Makefile in your app
