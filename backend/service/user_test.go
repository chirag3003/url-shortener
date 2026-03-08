package service

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/repository/mock"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"
)

func newTestUserService(t *testing.T) (*UserService, *mock.MockUserRepository) {
	ctrl := gomock.NewController(t)
	userRepo := mock.NewMockUserRepository(ctrl)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	svc := NewUserService(userRepo, log)
	return svc, userRepo
}

func TestGetByID_Success(t *testing.T) {
	svc, userRepo := newTestUserService(t)

	id := int64(1001)
	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(id, 10)).
		Return(&models.User{
			ID:      id,
			Name:    "Alice",
			Email:   "alice@example.com",
			PhoneNo: "+1234567890",
		}, nil)

	result, err := svc.GetByID(context.Background(), strconv.FormatInt(id, 10))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID != strconv.FormatInt(id, 10) {
		t.Fatalf("expected ID %s, got %s", strconv.FormatInt(id, 10), result.ID)
	}
	if result.Name != "Alice" {
		t.Fatalf("expected name Alice, got %s", result.Name)
	}
	if result.Email != "alice@example.com" {
		t.Fatalf("expected email alice@example.com, got %s", result.Email)
	}
	if result.PhoneNo != "+1234567890" {
		t.Fatalf("expected phoneNo +1234567890, got %s", result.PhoneNo)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	svc, userRepo := newTestUserService(t)

	id := int64(1002)
	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(id, 10)).
		Return(nil, nil)

	_, err := svc.GetByID(context.Background(), strconv.FormatInt(id, 10))
	if err == nil {
		t.Fatal("expected error for non-existent user, got nil")
	}

	var appErr *apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T: %v", err, err)
	}
	if appErr.Code != "NOT_FOUND" {
		t.Fatalf("expected code NOT_FOUND, got %s", appErr.Code)
	}
}

func TestGetByID_RepositoryError(t *testing.T) {
	svc, userRepo := newTestUserService(t)

	id := int64(1003)
	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(id, 10)).
		Return(nil, errors.New("db error"))

	_, err := svc.GetByID(context.Background(), strconv.FormatInt(id, 10))
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

func TestUpdate_Success(t *testing.T) {
	svc, userRepo := newTestUserService(t)

	id := int64(1001)
	existingUser := &models.User{
		ID:    id,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(id, 10)).
		Return(existingUser, nil)

	userRepo.EXPECT().
		UpdateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, u *models.User) error {
			if u.Name != "Alice Updated" {
				return errors.New("unexpected name")
			}
			if u.AvatarURL != "https://example.com/avatar.png" {
				return errors.New("unexpected avatar URL")
			}
			return nil
		})

	result, err := svc.Update(context.Background(), strconv.FormatInt(id, 10), "Alice Updated", "https://example.com/avatar.png")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Name != "Alice Updated" {
		t.Fatalf("expected name Alice Updated, got %s", result.Name)
	}
	if result.AvatarURL != "https://example.com/avatar.png" {
		t.Fatalf("expected avatar URL https://example.com/avatar.png, got %s", result.AvatarURL)
	}
}

func TestUpdate_PartialUpdate(t *testing.T) {
	svc, userRepo := newTestUserService(t)

	id := int64(1001)
	existingUser := &models.User{
		ID:    id,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	userRepo.EXPECT().
		GetUserByID(gomock.Any(), strconv.FormatInt(id, 10)).
		Return(existingUser, nil)

	userRepo.EXPECT().
		UpdateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	// Only update name
	result, err := svc.Update(context.Background(), strconv.FormatInt(id, 10), "New Name", "")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Name != "New Name" {
		t.Fatalf("expected name New Name, got %s", result.Name)
	}
}
