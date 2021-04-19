package handlers

import (
	"go_microservice/data"
	"net/http"
)

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable encoding list of products", http.StatusInternalServerError)
	}
}
