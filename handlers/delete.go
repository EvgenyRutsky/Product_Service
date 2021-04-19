package handlers

import (
	"github.com/gorilla/mux"
	"go_microservice/data"
	"net/http"
	"strconv"
)

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
