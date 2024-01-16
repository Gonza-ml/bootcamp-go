package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"time"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi"
)

type ProductHandlerDefault struct {
	service internal.ProductService
}

func NewProductHandler(hd internal.ProductService) *ProductHandlerDefault {
	return &ProductHandlerDefault{
		service: hd,
	}
}

// ProductJSON is the response for a product with json format
type ProductJSON struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

type ResponseBodyGetAll struct {
	Message string
	Data    []*ProductJSON
}

// GetAll get all products
func (hd *ProductHandlerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// process
		products := hd.service.GetAll()

		body := ResponseBodyGetAll{Message: "All products", Data: make([]*ProductJSON, 0, len(products))}
		for _, p := range products {
			body.Data = append(body.Data, &ProductJSON{
				Id:          p.Id,
				Name:        p.Name,
				Quantity:    p.Quantity,
				CodeValue:   p.CodeValue,
				IsPublished: p.IsPublished,
				Expiration:  p.Expiration,
				Price:       p.Price,
			})
		}

		// response
		response.JSON(w, http.StatusOK, body)
	}
}

type ResponseBodyGetById struct {
	Message string
	Data    *ProductJSON
}

// GetById get a product by id
func (hd *ProductHandlerDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// process
		product, err := hd.service.GetById(chi.URLParam(r, "id"))

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorInvalidId):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrorProductNotFound):
				response.Text(w, http.StatusNotFound, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		body := ResponseBodyGetById{Message: "Product", Data: &ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}}

		// response
		response.JSON(w, http.StatusOK, body)
	}
}

type ResponseBodySearch struct {
	Message string
	Data    []*ProductJSON
}

// SearchByPrice get a product by priceGt
func (hd *ProductHandlerDefault) SearchByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// process
		products, err := hd.service.SearchByPrice(r.URL.Query().Get("priceGt"))

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorInvalidPriceGt):
				response.Text(w, http.StatusBadRequest, err.Error())
			default:
				response.Text(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		body := ResponseBodySearch{Message: "Product", Data: make([]*ProductJSON, 0, len(products))}
		for _, p := range products {
			body.Data = append(body.Data, &ProductJSON{
				Id:          p.Id,
				Name:        p.Name,
				Quantity:    p.Quantity,
				CodeValue:   p.CodeValue,
				IsPublished: p.IsPublished,
				Expiration:  p.Expiration,
				Price:       p.Price,
			})
		}

		// response
		response.JSON(w, http.StatusOK, body)
	}
}

// RequestBodyProductCreate is the request for a product with json format
type RequestBodyProductCreate struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// ResponseBodyProductCreate is the response for a product with json format
type ResponseBodyProductCreate struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

type ResponseBodyCreate struct {
	Message string
	Data    *ResponseBodyProductCreate
}

func (hd *ProductHandlerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := &RequestBodyProductCreate{}
		err := request.JSON(r, &reqBody)
		if err != nil {
			// switch {
			// case errors.Is(err, request.ErrRequestContentTypeNotJSON):
			// 	response.Text(w, http.StatusBadRequest, request.ErrRequestContentTypeNotJSON.Error())
			// case errors.Is(err, request.ErrRequestJSONInvalid):
			// 	response.Text(w, http.StatusBadRequest, request.ErrRequestJSONInvalid.Error())
			// }
			response.Text(w, http.StatusBadRequest, err.Error())
			return
		}
		exp, err := time.Parse("02/01/2006", reqBody.Expiration)
		if err != nil {
			response.Text(w, http.StatusBadRequest, internal.ErrorInvalidExpiration.Error())
			return
		}
		// deserialize (refactorizar)
		product := &internal.Product{
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  exp,
			Price:       reqBody.Price,
		}

		prod, err := hd.service.Create(product)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorInvalidProduct):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrorInvalidCodeValue):
				response.Text(w, http.StatusBadRequest, "Invalid Product: codeValue")
			default:
				response.Text(w, http.StatusInternalServerError, "Internal server error")
			}
			return
		}

		body := &ResponseBodyCreate{
			Message: "Product create",
			Data: &ResponseBodyProductCreate{
				Id:          prod.Id,
				Name:        prod.Name,
				Quantity:    prod.Quantity,
				CodeValue:   prod.CodeValue,
				IsPublished: prod.IsPublished,
				Expiration:  prod.Expiration.Format("2006-01-02"),
				Price:       prod.Price,
			},
		}
		response.JSON(w, http.StatusOK, body)
	}
}
