package validate_test

import (
	"testing"

	"github.com/chirag3003/go-backend-template/pkg/validate"
)

type testStruct struct {
	Name     string `validate:"required,min=2"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func TestStruct_Valid(t *testing.T) {
	s := testStruct{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	if err := validate.Struct(s); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestStruct_MissingRequired(t *testing.T) {
	s := testStruct{
		Name:     "",
		Email:    "john@example.com",
		Password: "password123",
	}

	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error for missing name")
	}
}

func TestStruct_InvalidEmail(t *testing.T) {
	s := testStruct{
		Name:     "John",
		Email:    "not-an-email",
		Password: "password123",
	}

	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error for invalid email")
	}
}

func TestStruct_ShortPassword(t *testing.T) {
	s := testStruct{
		Name:     "John",
		Email:    "john@example.com",
		Password: "short",
	}

	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error for short password")
	}
}

func TestStruct_MultipleErrors(t *testing.T) {
	s := testStruct{
		Name:     "",
		Email:    "not-an-email",
		Password: "short",
	}

	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error for multiple fields")
	}
}
