package controller

import (
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
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
