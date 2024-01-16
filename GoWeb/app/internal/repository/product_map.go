package repository

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ProductMap is a struct that represent a Product repository
type ProductMap struct {
	Products map[int]internal.Product
	LastId   int
}

// ProductAttributesJSON is the body for a product in the json file
type ProductAttributesJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func NewProductMap(m map[int]internal.Product, lastId int) *ProductMap {
	return &ProductMap{
		Products: m,
		LastId:   lastId,
	}
}

// LoadProducts load products with a json file and return a products map
func LoadProductsJSON(fileName string) (productsMap map[int]internal.Product, err error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, internal.ErrorLoadProducts
	}
	defer file.Close()

	// Read file and decode json
	var products []ProductAttributesJSON
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return
	}

	// Serialize products
	productsMap = make(map[int]internal.Product)
	for _, p := range products {
		exp, err := time.Parse("02/01/2006", p.Expiration)
		if err != nil {
			return nil, err
		}
		productsMap[p.Id] = internal.Product{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  exp,
			Price:       p.Price,
		}
	}

	return
}

func (p *ProductMap) GetAll() map[int]internal.Product {
	return p.Products
}

func (p *ProductMap) GetById(id int) (product internal.Product, err error) {
	product, ok := p.Products[id]
	if !ok {
		err = internal.ErrorProductNotFound
	}
	return
}

func (p *ProductMap) GetByCodeValue(code string) (product internal.Product, err error) {
	for _, prod := range p.Products {
		if prod.CodeValue == code {
			fmt.Println(prod)
			product = prod
			return
		}
	}
	return
}

func (p *ProductMap) SearchByPrice(price float64) (products []internal.Product, err error) {
	for _, p := range p.Products {
		if p.Price > price {
			products = append(products, p)
		}
	}
	return
}

func (p *ProductMap) Create(product *internal.Product) (prod internal.Product, err error) {
	p.LastId++
	prod = internal.Product{
		Id:          p.LastId,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
	// Es mejor hacer una copia atributo por atributo para evitar shallowcopy
	p.Products[p.LastId] = prod
	return
}
