package validation

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

type ValidationError struct {
	Field string
	Tag   string
	Value string
}

func Init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
