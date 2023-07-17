package main

import (
	"ginTraining/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Creating object for controller
	productController := controllers.ProductController{}
	authcontroller := controllers.AuthController{}

	//Specifically for product routes
	productRoutes := r.Group("/product")

	productRoutes.Use() // We will be attaching some middlewares
	{
		productRoutes.POST("/", productController.AddProduct)

	}

	// Authroutes

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("signup", authcontroller.Signup)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
