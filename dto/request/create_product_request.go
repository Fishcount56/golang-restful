package request

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	Name string `validate:"required,min=3,max=255" json:"product_name"`
	Type string `validate:"required,min=3,max=255" json:"product_type"`
}

var Validate = validator.New()
