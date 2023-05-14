package gfun

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateStruct(d any) error {
	if validate == nil {
		validate = validator.New()
	}
	return validate.Struct(d)
}
