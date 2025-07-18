package controllers

import (
	"go-rest-api/models"
	"go-rest-api/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecases.ProductUsecase
}

func NewProductController(usecases usecases.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecases,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product models.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("cdproduct")

	if id == "" {
		response := models.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := models.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("cdproduct")

	if id == "" {
		response := models.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DeleteProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := models.Response{
		Message: "Product deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("cdproduct")

	if id == "" {
		response := models.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "Invalid ID, must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product models.Product
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.productUseCase.UpdateProduct(productId, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := models.Response{
		Message: "Product updated successfully",
	}
	ctx.JSON(http.StatusOK, response)
}