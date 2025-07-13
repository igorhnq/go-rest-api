package controllers

import (
	"go-rest-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecases.ProductUsecase
}

func NewProductController(usecases usecases.ProductUsecase) productController {
	return productController{
		productUseCase: usecases,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}