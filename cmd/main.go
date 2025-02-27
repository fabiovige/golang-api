package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// camada repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada usecase
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	// camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:productId", ProductController.GetProductById)
	server.PUT("/products/:productId", ProductController.UpdateProduct)
	server.DELETE("/products/:productId", ProductController.DeleteProduct)

	server.Run(":8000")
}
