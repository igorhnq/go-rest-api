package product

import (
	"go-rest-api/internal/domain/entities"
	"go-rest-api/internal/domain/repositories"
)

type ProductUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUseCase(repository repositories.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]entities.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product entities.Product) (entities.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return entities.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(cdproduct int) (*entities.Product, error) {
	product, err := pu.repository.GetProductById(cdproduct)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(cdproduct int) error {
	return pu.repository.DeleteProductById(cdproduct)
}

func (pu *ProductUsecase) UpdateProduct(cdproduct int, product entities.Product) error {
	return pu.repository.UpdateProduct(cdproduct, product)
}
