package controller

import (
	"strconv"

	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// AnalyticsController handles analytics endpoints.
type AnalyticsController struct {
	analyticsService *service.AnalyticsService
}

// NewAnalyticsController creates a new AnalyticsController.
func NewAnalyticsController(analyticsService *service.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{analyticsService: analyticsService}
}

// Summary handles GET /api/v1/links/:id/analytics/summary.
func (c *AnalyticsController) Summary(ctx fiber.Ctx) error {
	userID, linkID, err := parseAnalyticsIDs(ctx)
	if err != nil {
		return err
	}

	res, svcErr := c.analyticsService.GetSummary(ctx.Context(), userID, linkID)
	if svcErr != nil {
		return svcErr
	}
	return ctx.JSON(response.OK(res))
}

// TimeSeries handles GET /api/v1/links/:id/analytics/timeseries.
func (c *AnalyticsController) TimeSeries(ctx fiber.Ctx) error {
	userID, linkID, err := parseAnalyticsIDs(ctx)
	if err != nil {
		return err
	}

	res, svcErr := c.analyticsService.GetTimeSeries(ctx.Context(), userID, linkID, ctx.Query("window"))
	if svcErr != nil {
		return svcErr
	}
	return ctx.JSON(response.OK(res))
}

// Breakdown handles GET /api/v1/links/:id/analytics/:kind.
func (c *AnalyticsController) Breakdown(ctx fiber.Ctx) error {
	userID, linkID, err := parseAnalyticsIDs(ctx)
	if err != nil {
		return err
	}

	res, svcErr := c.analyticsService.GetBreakdown(ctx.Context(), userID, linkID, ctx.Params("kind"))
	if svcErr != nil {
		return svcErr
	}
	return ctx.JSON(response.OK(res))
}

func parseAnalyticsIDs(ctx fiber.Ctx) (int64, int64, error) {
	rawUserID, ok := ctx.Locals("userID").(string)
	if !ok || rawUserID == "" {
		return 0, 0, apperror.Unauthorized("user not authenticated")
	}
	userID, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		return 0, 0, apperror.Unauthorized("invalid authenticated user id")
	}

	linkID, parseErr := service.ParseID(ctx.Params("id"))
	if parseErr != nil {
		return 0, 0, parseErr
	}

	return userID, linkID, nil
}
