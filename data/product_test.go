package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{
		Name: "Test Name",
		Price: 1.0,
		SKU: "frs-asd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
