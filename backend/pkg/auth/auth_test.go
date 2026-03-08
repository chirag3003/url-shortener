package auth_test

import (
	"testing"
	"time"

	"github.com/chirag3003/go-backend-template/pkg/auth"
)

func TestHashPassword(t *testing.T) {
	password := "securepassword123"

	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hash == "" {
		t.Fatal("HashPassword returned empty hash")
	}

	if hash == password {
		t.Fatal("hash should not equal the plaintext password")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "securepassword123"

	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if !auth.VerifyPassword(password, hash) {
		t.Fatal("VerifyPassword should return true for correct password")
	}

	if auth.VerifyPassword("wrongpassword", hash) {
		t.Fatal("VerifyPassword should return false for incorrect password")
	}
}

func TestHashPasswordUniqueness(t *testing.T) {
	password := "securepassword123"

	hash1, _ := auth.HashPassword(password)
	hash2, _ := auth.HashPassword(password)

	if hash1 == hash2 {
		t.Fatal("two hashes of the same password should differ (due to salt)")
	}
}

func TestJWTService_GenerateAndParse(t *testing.T) {
	svc := auth.NewJWTService("test-secret", 1*time.Hour)

	token, err := svc.GenerateToken("user123", "John Doe", "john@example.com", "1234567890")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	if token == "" {
		t.Fatal("GenerateToken returned empty token")
	}

	claims, err := svc.ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken returned error: %v", err)
	}

	if claims.UserID != "user123" {
		t.Errorf("expected UserID 'user123', got '%s'", claims.UserID)
	}
	if claims.Name != "John Doe" {
		t.Errorf("expected Name 'John Doe', got '%s'", claims.Name)
	}
	if claims.Email != "john@example.com" {
		t.Errorf("expected Email 'john@example.com', got '%s'", claims.Email)
	}
	if claims.PhoneNo != "1234567890" {
		t.Errorf("expected PhoneNo '1234567890', got '%s'", claims.PhoneNo)
	}
}

func TestJWTService_InvalidToken(t *testing.T) {
	svc := auth.NewJWTService("test-secret", 1*time.Hour)

	_, err := svc.ParseToken("invalid-token")
	if err == nil {
		t.Fatal("ParseToken should return error for invalid token")
	}
}

func TestJWTService_WrongSecret(t *testing.T) {
	svc1 := auth.NewJWTService("secret-1", 1*time.Hour)
	svc2 := auth.NewJWTService("secret-2", 1*time.Hour)

	token, err := svc1.GenerateToken("user123", "John", "john@example.com", "")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	_, err = svc2.ParseToken(token)
	if err == nil {
		t.Fatal("ParseToken should return error for token signed with different secret")
	}
}

func TestJWTService_ExpiredToken(t *testing.T) {
	// Create a service with a negative expiration (already expired)
	svc := auth.NewJWTService("test-secret", -1*time.Hour)

	token, err := svc.GenerateToken("user123", "John", "john@example.com", "")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	_, err = svc.ParseToken(token)
	if err == nil {
		t.Fatal("ParseToken should return error for expired token")
	}
}
