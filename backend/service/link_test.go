package service

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/repository/mock"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"
)

func TestCreateLink_Parsing(t *testing.T) {
	ctrl := gomock.NewController(t)
	linkRepo := mock.NewMockLinkRepository(ctrl)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	svc := NewLinkService(linkRepo, "http://localhost:5000", log)

	userID := int64(1)

	tests := []struct {
		name         string
		expiresAt    string
		redirectType int16
		expectedType int16
	}{
		{"RFC3339", "2030-01-02T15:04:05Z", 301, 301},
		{"ISO Date Only", "2030-01-02", 0, 302}, // default
		{"ISO with millis", "2030-01-02T15:04:05.000Z", 301, 301},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			linkRepo.EXPECT().AliasExists(gomock.Any(), gomock.Any()).Return(false, nil)
			linkRepo.EXPECT().CreateLink(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, l *models.Link) error {
				if l.RedirectType != tc.expectedType {
					t.Errorf("expected redirect type %d, got %d", tc.expectedType, l.RedirectType)
				}
				if tc.expiresAt != "" && l.ExpiresAt == nil {
					t.Error("expected expiresAt to be parsed, got nil")
				}
				if l.UserID == nil || *l.UserID != userID {
					t.Errorf("expected userID %d, got %v", userID, l.UserID)
				}
				return nil
			})
			linkRepo.EXPECT().GetLinkByID(gomock.Any(), gomock.Any()).Return(&models.Link{
				ID:        123,
				ShortCode: "abc",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

			req := &request.CreateLinkRequest{
				LongURL:      "https://google.com",
				ExpiresAt:    tc.expiresAt,
				RedirectType: tc.redirectType,
			}
			_, err := svc.Create(context.Background(), userID, req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestCreateLink_InvalidDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	linkRepo := mock.NewMockLinkRepository(ctrl)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	svc := NewLinkService(linkRepo, "http://localhost:5000", log)

	userID := int64(1)
	req := &request.CreateLinkRequest{
		LongURL:   "https://google.com",
		ExpiresAt: "invalid-date",
	}

	linkRepo.EXPECT().AliasExists(gomock.Any(), gomock.Any()).Return(false, nil)

	_, err := svc.Create(context.Background(), userID, req)
	if err == nil {
		t.Fatal("expected error for invalid date, got nil")
	}
}
