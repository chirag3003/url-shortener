package controller

import (
	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// AuthController handles authentication HTTP requests.
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController creates a new AuthController.
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login handles POST /api/v1/auth/login
func (c *AuthController) Login(ctx fiber.Ctx) error {
	var req request.LoginRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		return apperror.BadRequest("invalid request body")
	}

	result, err := c.authService.Login(ctx.Context(), &req)
	if err != nil {
		return err // handled by centralized error handler
	}

	return ctx.Status(fiber.StatusOK).JSON(response.OK(result))
}

// Register handles POST /api/v1/auth/register
func (c *AuthController) Register(ctx fiber.Ctx) error {
	var req request.RegisterRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		return apperror.BadRequest("invalid request body")
	}

	if err := c.authService.Register(ctx.Context(), &req); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.OK(
		response.MessageResponse{Message: "registration successful"},
	))
}
