package controller

import (
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// RedirectController handles public short-code redirects.
type RedirectController struct {
	linkService      *service.LinkService
	analyticsService *service.AnalyticsService
}

// NewRedirectController creates a new RedirectController.
func NewRedirectController(linkService *service.LinkService, analyticsService *service.AnalyticsService) *RedirectController {
	return &RedirectController{
		linkService:      linkService,
		analyticsService: analyticsService,
	}
}

// Redirect handles GET /:code.
func (c *RedirectController) Redirect(ctx fiber.Ctx) error {
	code := ctx.Params("code")
	link, err := c.linkService.ResolveForRedirect(ctx.Context(), code)
	if err != nil {
		return err
	}

	go c.analyticsService.RecordClick(
		ctx.Context(),
		link.ID,
		ctx.IP(),
		ctx.Get("User-Agent"),
		ctx.Get("Referer"),
	)

	if link.RedirectType == 301 {
		return ctx.Redirect().Status(fiber.StatusMovedPermanently).To(link.LongURL)
	}
	return ctx.Redirect().Status(fiber.StatusFound).To(link.LongURL)
}
