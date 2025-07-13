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

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO products" +
		"(nmproduct, vlprice)" +
		"VALUES ($1, $2) RETURNING cdproduct")
	if err != nil {
		fmt.Println("Error preparing product: ", err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println("Error creating product: ", err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductById(cdproduct int) (*models.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM products WHERE cdproduct = $1")
	if err != nil {
		fmt.Println("Error preparing product: ", err)
		return nil, err
	}

	var product models.Product

	err = query.QueryRow(cdproduct).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if (err == sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}