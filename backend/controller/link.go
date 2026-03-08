package controller

import (
	"strconv"

	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// LinkController handles short-link endpoints.
type LinkController struct {
	linkService *service.LinkService
}

// NewLinkController creates a new LinkController.
func NewLinkController(linkService *service.LinkService) *LinkController {
	return &LinkController{linkService: linkService}
}

// Create handles POST /api/v1/links.
func (c *LinkController) Create(ctx fiber.Ctx) error {
	var req request.CreateLinkRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		return apperror.BadRequest("invalid request body")
	}

	userID, err := parseUserID(ctx)
	if err != nil {
		return err
	}

	res, err := c.linkService.Create(ctx.Context(), userID, &req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.OK(res))
}

// GetByID handles GET /api/v1/links/:id.
func (c *LinkController) GetByID(ctx fiber.Ctx) error {
	userID, linkID, err := parseUserAndLinkID(ctx)
	if err != nil {
		return err
	}

	res, svcErr := c.linkService.GetByID(ctx.Context(), userID, linkID)
	if svcErr != nil {
		return svcErr
	}

	return ctx.JSON(response.OK(res))
}

// List handles GET /api/v1/links.
func (c *LinkController) List(ctx fiber.Ctx) error {
	userID, err := parseUserID(ctx)
	if err != nil {
		return err
	}

	req := request.ListLinksRequest{}
	if bindErr := ctx.Bind().Query(&req); bindErr != nil {
		return apperror.BadRequest("invalid query params")
	}

	res, svcErr := c.linkService.ListByUser(ctx.Context(), userID, &req)
	if svcErr != nil {
		return svcErr
	}

	return ctx.JSON(response.OK(res))
}

// Update handles PATCH /api/v1/links/:id.
func (c *LinkController) Update(ctx fiber.Ctx) error {
	userID, linkID, err := parseUserAndLinkID(ctx)
	if err != nil {
		return err
	}

	var req request.UpdateLinkRequest
	if bindErr := ctx.Bind().JSON(&req); bindErr != nil {
		return apperror.BadRequest("invalid request body")
	}

	res, svcErr := c.linkService.Update(ctx.Context(), userID, linkID, &req)
	if svcErr != nil {
		return svcErr
	}

	return ctx.JSON(response.OK(res))
}

// Delete handles DELETE /api/v1/links/:id.
func (c *LinkController) Delete(ctx fiber.Ctx) error {
	userID, linkID, err := parseUserAndLinkID(ctx)
	if err != nil {
		return err
	}

	if svcErr := c.linkService.Delete(ctx.Context(), userID, linkID); svcErr != nil {
		return svcErr
	}

	return ctx.JSON(response.OK(response.MessageResponse{Message: "link deleted"}))
}

// AliasAvailability handles GET /api/v1/links/alias-availability.
func (c *LinkController) AliasAvailability(ctx fiber.Ctx) error {
	alias := ctx.Query("alias")
	res, err := c.linkService.CheckAliasAvailability(ctx.Context(), alias)
	if err != nil {
		return err
	}
	return ctx.JSON(response.OK(res))
}

func parseUserID(ctx fiber.Ctx) (int64, error) {
	raw, ok := ctx.Locals("userID").(string)
	if !ok || raw == "" {
		return 0, apperror.Unauthorized("user not authenticated")
	}
	parsed, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, apperror.Unauthorized("invalid authenticated user id")
	}
	return parsed, nil
}

func parseUserAndLinkID(ctx fiber.Ctx) (int64, int64, error) {
	userID, err := parseUserID(ctx)
	if err != nil {
		return 0, 0, err
	}

	linkID, parseErr := service.ParseID(ctx.Params("id"))
	if parseErr != nil {
		return 0, 0, parseErr
	}

	return userID, linkID, nil
}
