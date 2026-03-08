package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/repository/mock"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"
)

type stubIDGen struct {
	next int64
}

func (s *stubIDGen) NewID() (int64, error) {
	id := s.next
	s.next++
	return id, nil
}

var _ idgen.Generator = (*stubIDGen)(nil)

func newTestAuthService(t *testing.T) (*AuthService, *mock.MockUserRepository) {
	ctrl := gomock.NewController(t)
	userRepo := mock.NewMockUserRepository(ctrl)
	jwtService := auth.NewJWTService("test-secret", 1*time.Hour)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	idgen.SetDefault(&stubIDGen{next: 1000})
	svc := NewAuthService(userRepo, jwtService, log)
	return svc, userRepo
}

// --- Register Tests ---

func TestRegister_Success(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "alice@example.com").
		Return(nil, nil)
	userRepo.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "securepass123",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "alice@example.com").
		Return(&models.User{
			ID:    int64(1),
			Name:  "Alice",
			Email: "alice@example.com",
		}, nil)

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "Alice 2",
		Email:    "alice@example.com",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "USER_ALREADY_EXISTS" {
		t.Fatalf("expected code USER_ALREADY_EXISTS, got %s", appErr.Code)
	}
}

func TestRegister_ValidationError_MissingName(t *testing.T) {
	svc, _ := newTestAuthService(t)

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "",
		Email:    "alice@example.com",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected code VALIDATION_ERROR, got %s", appErr.Code)
	}
}

func TestRegister_ValidationError_ShortPassword(t *testing.T) {
	svc, _ := newTestAuthService(t)

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "short",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected code VALIDATION_ERROR, got %s", appErr.Code)
	}
}

func TestRegister_ValidationError_InvalidEmail(t *testing.T) {
	svc, _ := newTestAuthService(t)

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "Alice",
		Email:    "not-an-email",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected code VALIDATION_ERROR, got %s", appErr.Code)
	}
}

func TestRegister_RepositoryError(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "alice@example.com").
		Return(nil, nil)
	userRepo.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(errors.New("db connection lost"))

	err := svc.Register(context.Background(), &request.RegisterRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "INTERNAL_ERROR" {
		t.Fatalf("expected code INTERNAL_ERROR, got %s", appErr.Code)
	}
}

// --- Login Tests ---

func TestLogin_Success(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	hash, _ := auth.HashPassword("securepass123")
	userID := int64(987654)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "bob@example.com").
		Return(&models.User{
			ID:    userID,
			Name:  "Bob",
			Email: "bob@example.com",
			Hash:  hash,
		}, nil)

	resp, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "bob@example.com",
		Password: "securepass123",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.Token == "" {
		t.Fatal("expected non-empty token")
	}
	if resp.User.Email != "bob@example.com" {
		t.Fatalf("expected email bob@example.com, got %s", resp.User.Email)
	}
	if resp.User.Name != "Bob" {
		t.Fatalf("expected name Bob, got %s", resp.User.Name)
	}
	if resp.User.ID != strconv.FormatInt(userID, 10) {
		t.Fatalf("expected ID %s, got %s", strconv.FormatInt(userID, 10), resp.User.ID)
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	hash, _ := auth.HashPassword("securepass123")

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "bob@example.com").
		Return(&models.User{
			ID:    int64(2),
			Name:  "Bob",
			Email: "bob@example.com",
			Hash:  hash,
		}, nil)

	_, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "bob@example.com",
		Password: "wrongpassword",
	})
	if err == nil {
		t.Fatal("expected error for invalid password, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "INVALID_CREDENTIALS" {
		t.Fatalf("expected code INVALID_CREDENTIALS, got %s", appErr.Code)
	}
}

func TestLogin_NonExistentUser(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "nobody@example.com").
		Return(nil, nil)

	_, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "nobody@example.com",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected error for non-existent user, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "INVALID_CREDENTIALS" {
		t.Fatalf("expected code INVALID_CREDENTIALS, got %s", appErr.Code)
	}
}

func TestLogin_ValidationError_MissingEmail(t *testing.T) {
	svc, _ := newTestAuthService(t)

	_, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected code VALIDATION_ERROR, got %s", appErr.Code)
	}
}

func TestLogin_ValidationError_ShortPassword(t *testing.T) {
	svc, _ := newTestAuthService(t)

	_, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "bob@example.com",
		Password: "short",
	})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "VALIDATION_ERROR" {
		t.Fatalf("expected code VALIDATION_ERROR, got %s", appErr.Code)
	}
}

func TestLogin_RepositoryError(t *testing.T) {
	svc, userRepo := newTestAuthService(t)

	userRepo.EXPECT().
		GetUserByEmail(gomock.Any(), "bob@example.com").
		Return(nil, errors.New("db read error"))

	_, err := svc.Login(context.Background(), &request.LoginRequest{
		Email:    "bob@example.com",
		Password: "securepass123",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "INTERNAL_ERROR" {
		t.Fatalf("expected code INTERNAL_ERROR, got %s", appErr.Code)
	}
}
