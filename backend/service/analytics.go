package service

import (
	"context"
	"strconv"
	"time"

	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/rs/zerolog"
	"github.com/ua-parser/uap-go/uaparser"
)

var uaParser *uaparser.Parser

func init() {
	uaParser = uaparser.NewFromSaved()
}

// AnalyticsService handles click ingestion and analytics queries.
type AnalyticsService struct {
	clickRepo repository.ClickRepository
	linkRepo  repository.LinkRepository
	log       zerolog.Logger
}

// NewAnalyticsService creates a new AnalyticsService.
func NewAnalyticsService(clickRepo repository.ClickRepository, linkRepo repository.LinkRepository, log zerolog.Logger) *AnalyticsService {
	return &AnalyticsService{
		clickRepo: clickRepo,
		linkRepo:  linkRepo,
		log:       log.With().Str("service", "analytics").Logger(),
	}
}

// RecordClick stores one click event for redirect telemetry.
func (s *AnalyticsService) RecordClick(ctx context.Context, linkID int64, ip string, userAgent string, referrer string) {
	clickID, err := idgen.NewID()
	if err != nil {
		s.log.Error().Err(err).Int64("linkID", linkID).Msg("failed to generate click id")
		return
	}

	client := uaParser.Parse(userAgent)

	click := &models.ClickEvent{
		ID:         clickID,
		LinkID:     linkID,
		IPAddress:  ip,
		UserAgent:  userAgent,
		Referrer:   referrer,
		Browser:    client.UserAgent.Family,
		DeviceType: client.Device.Family,
	}

	if err := s.clickRepo.CreateClick(ctx, click); err != nil {
		s.log.Error().Err(err).Int64("linkID", linkID).Msg("failed to store click event")
	}
}

// GetSummary returns top-level analytics metrics for a link.
func (s *AnalyticsService) GetSummary(ctx context.Context, userID int64, linkID int64) (*response.AnalyticsSummaryResponse, error) {
	if err := s.ensureLinkOwnership(ctx, userID, linkID); err != nil {
		return nil, err
	}

	total, unique, last24h, last7d, err := s.clickRepo.GetSummary(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	prevTotal, prevUnique, prev24h, prev7d, err := s.clickRepo.GetPreviousSummary(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	return &response.AnalyticsSummaryResponse{
		TotalClicks:           total,
		TotalClicksChange:     calculateChange(total, prevTotal),
		UniqueVisitors:        unique,
		UniqueVisitorsChange:  calculateChange(unique, prevUnique),
		ClicksLast24H:         last24h,
		ClicksLast24HChange:   calculateChange(last24h, prev24h),
		ClicksLast7Days:       last7d,
		ClicksLast7DaysChange: calculateChange(last7d, prev7d),
	}, nil
}

func calculateChange(current, previous int64) float64 {
	if previous == 0 {
		if current > 0 {
			return 100.0
		}
		return 0.0
	}
	return (float64(current-previous) / float64(previous)) * 100.0
}

// GetTimeSeries returns click counts in time buckets.
func (s *AnalyticsService) GetTimeSeries(ctx context.Context, userID int64, linkID int64, window string) ([]response.AnalyticsPoint, error) {
	if err := s.ensureLinkOwnership(ctx, userID, linkID); err != nil {
		return nil, err
	}

	if window == "" {
		window = "30d"
	}
	if window != "24h" && window != "7d" && window != "30d" {
		return nil, apperror.BadRequest("window must be one of 24h, 7d, 30d")
	}

	points, err := s.clickRepo.GetTimeSeries(ctx, linkID, window)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	res := make([]response.AnalyticsPoint, 0, len(points))
	for _, p := range points {
		res = append(res, response.AnalyticsPoint{
			Bucket: p.Bucket.Format(time.RFC3339),
			Clicks: p.Count,
		})
	}
	return res, nil
}

// GetBreakdown returns grouped analytics values for one dimension.
func (s *AnalyticsService) GetBreakdown(ctx context.Context, userID int64, linkID int64, kind string) ([]response.BreakdownItem, error) {
	if err := s.ensureLinkOwnership(ctx, userID, linkID); err != nil {
		return nil, err
	}

	fieldMap := map[string]string{
		"referrers": "referrer",
		"devices":   "device_type",
		"browsers":  "browser",
		"geography": "country",
	}
	field, ok := fieldMap[kind]
	if !ok {
		return nil, apperror.BadRequest("unsupported breakdown type")
	}

	rows, err := s.clickRepo.GetTopBreakdown(ctx, linkID, field, 15)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	result := make([]response.BreakdownItem, 0, len(rows))
	for _, row := range rows {
		result = append(result, response.BreakdownItem{Key: row.Key, Count: row.Count})
	}
	return result, nil
}

func (s *AnalyticsService) ensureLinkOwnership(ctx context.Context, userID int64, linkID int64) error {
	link, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return apperror.Internal(err)
	}
	if link == nil {
		return apperror.NotFound("link not found")
	}
	if link.UserID == nil || *link.UserID != userID {
		return apperror.ErrForbidden
	}
	return nil
}

// ParseID parses route ID strings.
func ParseID(id string) (int64, *apperror.AppError) {
	parsed, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, apperror.BadRequest("invalid id")
	}
	return parsed, nil
}
