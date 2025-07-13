package main

import (
	"go-rest-api/controllers"
	"go-rest-api/db"
	"go-rest-api/repositories"
	"go-rest-api/usecases"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repositories.NewProductRepository(dbConnection)

	ProductUseCase := usecases.NewProductUseCase(ProductRepository)

	ProductControler := controllers.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductControler.GetProducts)
	server.POST("/products", ProductControler.CreateProduct)

	server.Run(":8080")
}