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
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DeleteProductById(cdproduct int) error {
	query, err := pr.connection.Prepare("DELETE FROM products WHERE cdproduct = $1")
	if err != nil {
		fmt.Println("Error preparing delete product: ", err)
		return err
	}

	result, err := query.Exec(cdproduct)
	if err != nil {
		fmt.Println("Error deleting product: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected: ", err)
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	query.Close()
	return nil
}

func (pr *ProductRepository) UpdateProduct(cdproduct int, product models.Product) error {
	query, err := pr.connection.Prepare("UPDATE products SET nmproduct = $1, vlprice = $2 WHERE cdproduct = $3")
	if err != nil {
		fmt.Println("Error preparing update product: ", err)
		return err
	}

	result, err := query.Exec(product.Name, product.Price, cdproduct)
	if err != nil {
		fmt.Println("Error updating product: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected: ", err)
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	query.Close()
	return nil
}