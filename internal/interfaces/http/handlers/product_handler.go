package handlers

import (
	"go-rest-api/internal/domain/entities"
	"go-rest-api/internal/usecases/product"
	"go-rest-api/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase product.ProductUsecase
}

func NewProductHandler(usecase product.ProductUsecase) ProductHandler {
	return ProductHandler{
		productUseCase: usecase,
	}
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product entities.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("cdproduct")

	if id == "" {
		response := response.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := response.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		response := response.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductHandler) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("cdproduct")

	if id == "" {
		response := response.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := response.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := response.Response{
		Message: "Product deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("cdproduct")

	if id == "" {
		response := response.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := response.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product entities.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.productUseCase.UpdateProduct(productId, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := response.Response{
		Message: "Product updated successfully",
	}
	ctx.JSON(http.StatusOK, response)
}
