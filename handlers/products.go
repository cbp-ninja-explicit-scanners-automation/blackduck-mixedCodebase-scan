package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/amitramachandran/zero1/data"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProd{}).(*data.Product)
	data.AddProduct(*prod)
	p.l.Printf("Added product %#v", prod)
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ProductList := data.GetProducts()
	err := ProductList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error while encoding", http.StatusBadRequest)
	}

}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	prodId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProd{}).(*data.Product)
	err = data.UpdateProduct(*prod, prodId)
	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

type KeyProd struct{}

func (p *Product) ProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error while decoding", http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), KeyProd{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})

}
