package apperror

import (
	"fmt"
	"net/http"
)

// AppError represents a structured application error with an HTTP status code.
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError.
func New(code string, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// Wrap wraps an existing error into an AppError.
func Wrap(err error, code string, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
		Err:     err,
	}
}

// Pre-defined application errors.
var (
	ErrBadRequest         = New("BAD_REQUEST", "Invalid request", http.StatusBadRequest)
	ErrUnauthorized       = New("UNAUTHORIZED", "Unauthorized", http.StatusUnauthorized)
	ErrForbidden          = New("FORBIDDEN", "Forbidden", http.StatusForbidden)
	ErrNotFound           = New("NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrConflict           = New("CONFLICT", "Resource already exists", http.StatusConflict)
	ErrValidation         = New("VALIDATION_ERROR", "Validation failed", http.StatusUnprocessableEntity)
	ErrInternal           = New("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
	ErrInvalidCredentials = New("INVALID_CREDENTIALS", "Invalid email or password", http.StatusUnauthorized)
	ErrUserAlreadyExists  = New("USER_ALREADY_EXISTS", "A user with this email already exists", http.StatusConflict)
)

// BadRequest creates a bad request error with a custom message.
func BadRequest(message string) *AppError {
	return New("BAD_REQUEST", message, http.StatusBadRequest)
}

// NotFound creates a not found error with a custom message.
func NotFound(message string) *AppError {
	return New("NOT_FOUND", message, http.StatusNotFound)
}

// Internal wraps an internal error with a safe message for the client.
func Internal(err error) *AppError {
	return Wrap(err, "INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
}

// Unauthorized creates an unauthorized error with a custom message.
func Unauthorized(message string) *AppError {
	return New("UNAUTHORIZED", message, http.StatusUnauthorized)
}

// ValidationError creates a validation error with field-level details.
func ValidationError(message string) *AppError {
	return New("VALIDATION_ERROR", message, http.StatusUnprocessableEntity)
}
