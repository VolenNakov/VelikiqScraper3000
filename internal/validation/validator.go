package validation

import (
	"OlxScraper/internal/response"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// Validate validates a struct and returns validation errors
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// HandleValidationErrors converts validator.ValidationErrors to our custom format
func HandleValidationErrors(err error) []response.ValidationError {
	var validationErrors []response.ValidationError

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			validationErrors = append(validationErrors, response.ValidationError{
				Field: err.Field(),
				Error: translateError(err),
				Value: err.Value(),
			})
		}
	}

	return validationErrors
}

// translateError converts validator errors into human-readable messages
func translateError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "min":
		return "Must be at least " + err.Param() + " characters long"
	case "gte":
		return "Must be greater than or equal to " + err.Param()
	default:
		return "Invalid value"
	}
}
