package controller

import (
	"github.com/chirag3003/go-backend-template/service"
)

// Controllers aggregates all HTTP controllers.
type Controllers struct {
	Auth      *AuthController
	User      *UserController
	Media     *MediaController
	Link      *LinkController
	Redirect  *RedirectController
	Analytics *AnalyticsController
}

// NewControllers creates all controllers with the given services.
func NewControllers(
	authService *service.AuthService,
	userService *service.UserService,
	mediaService *service.MediaService,
	linkService *service.LinkService,
	analyticsService *service.AnalyticsService,
) *Controllers {
	return &Controllers{
		Auth:      NewAuthController(authService),
		User:      NewUserController(userService),
		Media:     NewMediaController(mediaService),
		Link:      NewLinkController(linkService),
		Redirect:  NewRedirectController(linkService, analyticsService),
		Analytics: NewAnalyticsController(analyticsService),
	}
}
