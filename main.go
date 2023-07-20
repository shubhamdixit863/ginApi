package main

import (
	"fmt"
	"ginTraining/internal/controllers"
	"ginTraining/internal/repository"
	"ginTraining/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MysqlConnect(uri string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected with Mysql")

	return db
}

func main() {

	/*
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/

	r := gin.Default()

	// Repository
	dsn := "admin:o9Uusjfn@tcp(mysql-135552-0.cloudclusters.net:17741)/mysqlGolang?charset=utf8mb4&parseTime=True&loc=Local"

	mysqlRepo := repository.MysqlAuthRepository{Db: MysqlConnect(dsn)}
	productRepo := repository.MysqlProductRepository{Db: MysqlConnect(dsn)}

	// Service
	svc := services.AuthService{AuthRepo: &mysqlRepo}
	pvd := services.ProductService{ProductRepo: &productRepo}

	// Creating object for controller
	productController := controllers.ProductController{ProductService: pvd}
	authcontroller := controllers.AuthController{AuthService: &svc}

	//Specifically for product routes
	productRoutes := r.Group("/product")

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	productRoutes.Use() // We will be attaching some middlewares
	{
		productRoutes.POST("/", productController.AddProduct)

	}

	// Authroutes

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", authcontroller.Signup)
		authRoutes.POST("/login", authcontroller.Login)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
