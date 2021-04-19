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
	"go_microservice/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}
var KeyProduct struct{}

func NewProduct (l *log.Logger) *Products {
	return &Products{l}
}

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
