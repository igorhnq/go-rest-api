package routes

import (
	"go-rest-api/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productHandler handlers.ProductHandler) {
	products := router.Group("/products")
	{
		products.GET("", productHandler.GetProducts)
		products.POST("", productHandler.CreateProduct)
		products.GET("/:cdproduct", productHandler.GetProductById)
		products.PUT("/:cdproduct", productHandler.UpdateProduct)
		products.DELETE("/:cdproduct", productHandler.DeleteProductById)
	}
}
