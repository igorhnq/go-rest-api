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

func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(cdproduct int) (*models.Product, error) {
	product, err := pu.repository.GetProductById(cdproduct)

	if err != nil {
		return nil, err
	}

	return product, nil	
}