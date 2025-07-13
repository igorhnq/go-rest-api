package repositories

import (
	"database/sql"
	"fmt"
	"go-rest-api/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (p *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT * FROM products"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println("Error querying products: ", err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)
		if err != nil {
			fmt.Println("Error scanning product: ", err)
			return []models.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
