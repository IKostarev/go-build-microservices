package handlers

import (
	"go-build-microservices-product-api/storage"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	listProduct := storage.GetProducts()

	err := listProduct.ToJSON(w)
	if err != nil {
		http.Error(w, "func ToJSON return error\n", http.StatusInternalServerError)
	}
}
