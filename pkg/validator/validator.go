package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator wraps go-playground validator
type Validator struct {
	validate *validator.Validate
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	validate := validator.New()

	// Use JSON tag names for field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom validations here
	// validate.RegisterValidation("custom_tag", customValidationFunc)

	return &Validator{validate: validate}
}

// Validate validates a struct
func (v *Validator) Validate(data interface{}) []ValidationError {
	var errors []ValidationError

	err := v.validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Tag:     err.Tag(),
				Value:   err.Param(),
				Message: getErrorMessage(err),
			})
		}
	}

	return errors
}

// getErrorMessage returns a human-readable error message
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	case "eqfield":
		return fe.Field() + " must match " + fe.Param()
	case "url":
		return fe.Field() + " must be a valid URL"
	default:
		return fe.Field() + " is invalid"
	}
}
