package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productController controllers.ProductController) {
	SetupProductRoutes(router, productController)
}