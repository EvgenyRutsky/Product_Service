package handlers

import (
	"github.com/gorilla/mux"
	"go_microservice/data"
	"net/http"
	"strconv"
)

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
