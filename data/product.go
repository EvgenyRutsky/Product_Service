package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
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

func (p *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lastP := productList[len(productList) - 1]
	return lastP.ID + 1
}

type Products []*Product

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

func GetProducts()Products {
	return productList
}

var ErrProductNotFound error = fmt.Errorf("product not found")

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

func findProduct(id int) (*Product, int,  error) {
	for ind, val := range productList {
		if val.ID == id {

			return val, ind, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func (p *Products) ToJSON (w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}



