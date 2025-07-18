package main

import (
	"go-rest-api/internal/domain/repositories"
	db "go-rest-api/internal/infrastructure/database"
	"go-rest-api/internal/interfaces/http/handlers"
	"go-rest-api/internal/interfaces/http/routes"
	"go-rest-api/internal/usecases/product"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repositories.NewProductRepository(dbConnection)
	productUseCase := product.NewProductUseCase(productRepository)
	productHandler := handlers.NewProductHandler(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetupRoutes(server, productHandler)

	server.Run(":8080")
}
