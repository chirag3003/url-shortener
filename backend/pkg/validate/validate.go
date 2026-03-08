package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator wraps the go-playground validator.
var v = validator.New()

// Struct validates a struct based on its validate tags.
// Returns a human-readable error message if validation fails.
func Struct(s interface{}) error {
	if err := v.Struct(s); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		var messages []string
		for _, e := range validationErrors {
			messages = append(messages, formatError(e))
		}
		return fmt.Errorf("%s", strings.Join(messages, "; "))
	}
	return nil
}

func formatError(e validator.FieldError) string {
	field := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
