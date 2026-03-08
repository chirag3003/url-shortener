package middleware

import (
	"strings"

	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/gofiber/fiber/v3"
)

// Auth returns a middleware that validates JWT tokens from the Authorization header.
func Auth(jwtService *auth.JWTService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return apperror.Unauthorized("missing authorization header")
		}

		// Support "Bearer <token>" format
		token := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		claims, err := jwtService.ParseToken(token)
		if err != nil {
			return apperror.Unauthorized("invalid or expired token")
		}

		ctx.Locals("userID", claims.UserID)
		ctx.Locals("userName", claims.Name)
		ctx.Locals("userEmail", claims.Email)
		ctx.Locals("userPhoneNo", claims.PhoneNo)
		ctx.Locals("authType", "jwt")
		return ctx.Next()
	}
}
