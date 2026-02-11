package main
import (
	"github.com/gin-gonic/gin"
	"GoApi/controller"
)

func main() {

	server := gin.Default()

	// Camada de Controller
	ProductController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}