package validate

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// Validator holds a structure for the go-playground
// validator module.
type Validator struct {
	validator *validator.Validate
}

// Validate implements the Echo#Validator handle
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// EnableValidator configures the Echo#Validator element
func EnableValidator(e *echo.Echo) {
	e.Validator = &Validator{validator: validator.New()}
}
