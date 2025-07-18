package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productController controllers.ProductController) {
	products := router.Group("/products")
	{
		products.GET("", productController.GetProducts)
		products.POST("", productController.CreateProduct)
		products.GET("/:cdproduct", productController.GetProductById)
		products.PUT("/:cdproduct", productController.UpdateProduct)
		products.DELETE("/:cdproduct", productController.DeleteProductById)
	}
}