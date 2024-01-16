package internal

import (
	"errors"
)

var (
	ErrorLoadProducts    = errors.New("Failed to open file")
	ErrorProductNotFound = errors.New("Product not found")
)

type ProductRepository interface {
	GetAll() map[int]Product
	GetById(id int) (Product, error)
	GetByCodeValue(code string) (Product, error)
	SearchByPrice(price float64) ([]Product, error)
	Create(product *Product) (Product, error)
}
