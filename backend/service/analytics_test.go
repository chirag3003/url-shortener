package service

import (
	"context"
	"os"
	"testing"

	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/repository/mock"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"
)

func TestGetSummary_CalculatesTrends(t *testing.T) {
	ctrl := gomock.NewController(t)
	clickRepo := mock.NewMockClickRepository(ctrl)
	linkRepo := mock.NewMockLinkRepository(ctrl)
	log := zerolog.New(os.Stderr).Level(zerolog.Disabled)
	svc := NewAnalyticsService(clickRepo, linkRepo, log)

	userID := int64(1)
	linkID := int64(100)

	linkRepo.EXPECT().GetLinkByID(gomock.Any(), linkID).Return(&models.Link{
		ID:     linkID,
		UserID: &userID,
	}, nil)

	// current: 150 total, 50 unique, 20 last24h, 100 last7d
	clickRepo.EXPECT().GetSummary(gomock.Any(), linkID).Return(int64(150), int64(50), int64(20), int64(100), nil)

	// previous: 100 total, 40 unique, 10 last24h, 80 last7d
	clickRepo.EXPECT().GetPreviousSummary(gomock.Any(), linkID).Return(int64(100), int64(40), int64(10), int64(80), nil)

	res, err := svc.GetSummary(context.Background(), userID, linkID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 150 vs 100 = +50%
	if res.TotalClicksChange != 50.0 {
		t.Errorf("expected 50%% change, got %f", res.TotalClicksChange)
	}
	// 50 vs 40 = +25%
	if res.UniqueVisitorsChange != 25.0 {
		t.Errorf("expected 25%% change, got %f", res.UniqueVisitorsChange)
	}
	// 20 vs 10 = +100%
	if res.ClicksLast24HChange != 100.0 {
		t.Errorf("expected 100%% change, got %f", res.ClicksLast24HChange)
	}
	// 100 vs 80 = +25%
	if res.ClicksLast7DaysChange != 25.0 {
		t.Errorf("expected 25%% change, got %f", res.ClicksLast7DaysChange)
	}
}

func TestCalculateChange(t *testing.T) {
	tests := []struct {
		current  int64
		previous int64
		expected float64
	}{
		{150, 100, 50.0},
		{50, 100, -50.0},
		{100, 100, 0.0},
		{100, 0, 100.0},
		{0, 0, 0.0},
		{0, 100, -100.0},
	}

	for _, tc := range tests {
		got := calculateChange(tc.current, tc.previous)
		if got != tc.expected {
			t.Errorf("calculateChange(%d, %d) = %f, expected %f", tc.current, tc.previous, got, tc.expected)
		}
	}
}
