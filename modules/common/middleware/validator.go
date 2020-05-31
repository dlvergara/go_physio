package common

import (
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator struct
type CustomValidator struct {
	validator *validator.Validate
}

// Validate request data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Validator custom
func Validator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
