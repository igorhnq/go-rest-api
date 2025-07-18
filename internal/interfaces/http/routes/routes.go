package routes

import (
	"go-rest-api/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler handlers.ProductHandler) {
	SetupProductRoutes(router, productHandler)
}
