// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//   - application/json
//
//  Produces:
//   - application/json
//  swagger:meta

package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go_microservice/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProduct (l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable encoding list of products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct).(*data.Product)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct (rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid product id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct).(*data.Product)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Error during the product update", http.StatusInternalServerError)
		return
	}
}

func (p *Products) DeleteProduct (rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Error parsing Product ID", http.StatusBadRequest)
	}

	p.l.Println("Handle Delete Product", id)

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Error during the product delete", http.StatusInternalServerError)
		return
	}

}

var KeyProduct struct{}
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode received JSON", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(
				rw,
				fmt.Sprintf("Error validation the product %s", err),
				http.StatusBadRequest,
				)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
