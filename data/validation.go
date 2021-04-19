package data

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

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
