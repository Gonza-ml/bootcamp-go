package service

import (
	"app/internal"
	"fmt"
	"strconv"
	"time"
)

type ProductServiceDefault struct {
	rp internal.ProductRepository
}

func NewProductServiceDefault(rep internal.ProductRepository) *ProductServiceDefault {
	return &ProductServiceDefault{
		rp: rep,
	}
}

// GetAll get all products
func (s *ProductServiceDefault) GetAll() map[int]internal.Product {
	return s.rp.GetAll()
}

// GetById get product by id
func (p *ProductServiceDefault) GetById(id string) (product internal.Product, err error) {
	pid, err := strconv.Atoi(id)
	if err != nil {
		return internal.Product{}, internal.ErrorInvalidId
	}
	product, err = p.rp.GetById(pid)
	if err != nil {
		switch err {
		case internal.ErrorProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrorProductNotFound)
		}
	}
	return
}

// GetByCodeValue get product by codeValue
func (p *ProductServiceDefault) GetByCodeValue(code string) (product internal.Product, err error) {
	product, err = p.rp.GetByCodeValue(code)
	if err != nil {
		return
	}
	return
}

// GetById get product by id
func (p *ProductServiceDefault) SearchByPrice(price string) (products []internal.Product, err error) {
	priceProduct, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return nil, internal.ErrorInvalidPriceGt
	}

	products, err = p.rp.SearchByPrice(priceProduct)
	if err != nil {
		return nil, err
	}
	return
}

func (p *ProductServiceDefault) Create(product *internal.Product) (prod internal.Product, err error) {
	//realizar validaciones
	if product.Name == "" {
		err = fmt.Errorf("%w: invalid name", internal.ErrorInvalidProduct)
		return
	}
	if product.Quantity < 0 {
		err = fmt.Errorf("%w: invalid quantity", internal.ErrorInvalidProduct)
		return
	}
	if product.CodeValue == "" {
		err = fmt.Errorf("%w: invalid code", internal.ErrorInvalidProduct)
		return
	}

	productCode, err := p.GetByCodeValue(prod.CodeValue)
	if err != nil {
		return
	}
	if productCode.CodeValue != "" {
		err = fmt.Errorf("%w: invalid code", internal.ErrorInvalidProduct)
		return
	}

	if !product.Expiration.Before(time.Now()) {
		err = fmt.Errorf("%w: invalid time", internal.ErrorInvalidProduct)
		return
	}
	if product.Price < 0 {
		err = fmt.Errorf("%w: invalid price", internal.ErrorInvalidProduct)
		return
	}

	prod, err = p.rp.Create(product)

	return

}
