package internal

import "errors"

var (
	ErrorInvalidId         = errors.New("Invalid id")
	ErrorInvalidCodeValue  = errors.New("Invalid code_value")
	ErrorInvalidPriceGt    = errors.New("Invalid priceGt")
	ErrorInvalidExpiration = errors.New("invalid date format. Must be yyyy-mm-dd")
	ErrorInvalidProduct    = errors.New("Invalid Product")
)

type ProductService interface {
	GetAll() map[int]Product
	GetById(id string) (Product, error)
	GetByCodeValue(code string) (Product, error)
	SearchByPrice(price string) ([]Product, error)
	Create(product *Product) (Product, error)
}
