package apperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/chirag3003/go-backend-template/pkg/apperror"
)

func TestNew(t *testing.T) {
	err := apperror.New("TEST_ERROR", "test message", http.StatusBadRequest)

	if err.Code != "TEST_ERROR" {
		t.Errorf("expected code 'TEST_ERROR', got '%s'", err.Code)
	}
	if err.Message != "test message" {
		t.Errorf("expected message 'test message', got '%s'", err.Message)
	}
	if err.Status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, err.Status)
	}
}

func TestAppError_Error(t *testing.T) {
	err := apperror.New("TEST", "test message", 400)
	expected := "TEST: test message"
	if err.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, err.Error())
	}
}

func TestAppError_ErrorWithWrapped(t *testing.T) {
	inner := errors.New("inner error")
	err := apperror.Wrap(inner, "TEST", "test message", 500)

	if !errors.Is(err, inner) {
		t.Error("Unwrap should return the inner error")
	}
}

func TestWrap(t *testing.T) {
	inner := errors.New("database connection failed")
	err := apperror.Wrap(inner, "DB_ERROR", "failed to connect", http.StatusInternalServerError)

	if err.Err != inner {
		t.Error("wrapped error should contain the original error")
	}

	if !errors.Is(err, inner) {
		t.Error("errors.Is should find the wrapped error")
	}
}

func TestBadRequest(t *testing.T) {
	err := apperror.BadRequest("invalid input")
	if err.Status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, err.Status)
	}
	if err.Message != "invalid input" {
		t.Errorf("expected message 'invalid input', got '%s'", err.Message)
	}
}

func TestInternal(t *testing.T) {
	inner := errors.New("something broke")
	err := apperror.Internal(inner)

	if err.Status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, err.Status)
	}
	if err.Err != inner {
		t.Error("Internal should wrap the provided error")
	}
	// Message should be safe for the client
	if err.Message != "Internal server error" {
		t.Errorf("expected safe message, got '%s'", err.Message)
	}
}

func TestPreDefinedErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      *apperror.AppError
		expected int
	}{
		{"ErrBadRequest", apperror.ErrBadRequest, http.StatusBadRequest},
		{"ErrUnauthorized", apperror.ErrUnauthorized, http.StatusUnauthorized},
		{"ErrForbidden", apperror.ErrForbidden, http.StatusForbidden},
		{"ErrNotFound", apperror.ErrNotFound, http.StatusNotFound},
		{"ErrConflict", apperror.ErrConflict, http.StatusConflict},
		{"ErrInternal", apperror.ErrInternal, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Status != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, tt.err.Status)
			}
		})
	}
}
