package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Run() (err error) {

	// load products to filename
	fileName := "./../products.json"
	productMap, err := repository.LoadProductsJSON(fileName)
	if err != nil {
		log.Println(err)
	}
	// init repository
	rp := repository.NewProductMap(productMap, len(productMap))

	// init service
	service := service.NewProductServiceDefault(rp)

	// init controller
	hd := handler.NewProductHandler(service)

	// init router
	rt := chi.NewRouter()

	// endpoints
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Get("/search", hd.SearchByPrice())
		r.Post("/", hd.Create())
	})

	// run server
	err = http.ListenAndServe(s.Addr, rt)
	return
}
