package controllers

import (
	"fmt"
	"ginTraining/internal/dtos"
	"ginTraining/internal/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type ProductController struct {
	ProductService services.ProductService
}

func (pr *ProductController) AddProduct(c *gin.Context) {
	// form data we will be dealing with
	// We will upload the file in local directory
	//fmt.Println("Name", c.PostForm("name"))
	//fmt.Println(c.PostForm("age"))
	file, err := c.FormFile("file")

	fmt.Println("Formdata--", c.PostForm("name"))
	fmt.Println("Formdata--", c.PostForm("description"))

	//fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}

	// We will upload it in the local directory
	uploadPath := ".././uploads"

	localFilePath := filepath.Join(uploadPath, filepath.Base(file.Filename))

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			return
		}
	}

	fmt.Println(localFilePath)
	err = c.SaveUploadedFile(file, localFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
			"error":   err.Error(),
		})
		return
	}
	// file upload in aws

	filePointer, err := os.Open(localFilePath)
	if err != nil {
		return
	}

	// aws session
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("your accesskey", "yoursecret", ""),
	})

	// you can upload the files in s3 ---

	uploader := s3manager.NewUploader(sess)

	upload, err := uploader.Upload(&s3manager.UploadInput{

		Body:   filePointer,
		Bucket: aws.String("recordingsstreamhbeon"),
	})
	if err != nil {
		return
	}

	fmt.Println(upload.Location) // file path of the uploaded file

	//  Service here ---->
	productDto := dtos.ProductDto{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Image:       localFilePath, //upload.Location)
	}

	err = pr.ProductService.SaveProduct(productDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
			"error":   err,
		})
		return
	}

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
