package usecases

import (
	"go-rest-api/models"
	"go-rest-api/repositories"
)

type ProductUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUseCase(repository repositories.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}