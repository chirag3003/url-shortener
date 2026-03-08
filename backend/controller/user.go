package controller

import (
	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/validate"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// UserController handles user HTTP requests.
type UserController struct {
	userService *service.UserService
}

// NewUserController creates a new UserController.
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetMe handles GET /api/v1/user/me
func (c *UserController) GetMe(ctx fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok || userID == "" {
		return apperror.Unauthorized("user not authenticated")
	}

	user, err := c.userService.GetByID(ctx.Context(), userID)
	if err != nil {
		return err
	}

	return ctx.JSON(response.OK(user))
}

// UpdateMe handles PATCH /api/v1/user/me
func (c *UserController) UpdateMe(ctx fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(string)
	if !ok || userID == "" {
		return apperror.Unauthorized("user not authenticated")
	}

	var req request.UpdateUserRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return apperror.BadRequest("invalid request body")
	}

	if err := validate.Struct(req); err != nil {
		return err
	}

	user, err := c.userService.Update(ctx.Context(), userID, req.Name, req.AvatarURL)
	if err != nil {
		return err
	}

	return ctx.JSON(response.OK(user))
}
