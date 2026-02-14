package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"GoApi/controller"
	"GoApi/usecase"
	"GoApi/repository"
	"GoApi/db"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	server := gin.Default()

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}
