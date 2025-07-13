package models

type Product struct {
	ID          int     `json:"cdproduct"`
	Name        string  `json:"nmproduct"`
	Price       float64 `json:"vlprice"`

}