package http

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ValidationErrorBag []ValidationError

func EncodeValidationErrorBag(bag ValidationErrorBag) string {
	jsonBody, _ := json.Marshal(bag)

	return string(jsonBody)
}

func FromValidationErrors(errors error) ValidationErrorBag {
	bag := make(ValidationErrorBag, 0)
	for _, err := range errors.(validator.ValidationErrors) {
		bag = append(bag, ValidationError{Message: err.Error(), Field: err.Field()})
	}

	return bag
}

func EncodeError(errors error) string {
	return EncodeValidationErrorBag(FromValidationErrors(errors))
}
