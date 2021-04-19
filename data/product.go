package data

import (
	"fmt"
	"time"
)

type Product struct {
	ID  int `json:"id"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price float32 `json:"price" validate:"gt=0"`
	SKU string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdateOn string `json:"-"`
	DeletedOn string `json:"-"`
}

var productList = []*Product {
	&Product{
		ID: 1,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		CreatedOn: time.Now().UTC().String(),
		UpdateOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Latte",
		Description: "Frothy ,milky coffee",
		Price: 2.45,
		CreatedOn: time.Now().UTC().String(),
		UpdateOn: time.Now().UTC().String(),
	},
}

var ErrProductNotFound error = fmt.Errorf("product not found")

type Products []*Product

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func GetProducts()Products {
	return productList
}

func UpdateProduct(id int, product *Product) error {
	_, ind,  err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[ind] = product
	return nil
}

func DeleteProduct(id int) error{
	_, ind, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:ind], productList[ind+1:]...)

	return nil

}

func getNextID() int {
	lastP := productList[len(productList) - 1]
	return lastP.ID + 1
}

func findProduct(id int) (*Product, int,  error) {
	for ind, val := range productList {
		if val.ID == id {

			return val, ind, nil
		}
	}

	return nil, -1, ErrProductNotFound
}





