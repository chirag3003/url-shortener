package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// RequestID adds a unique request ID to each request context and response header.
func RequestID() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		id := ctx.Get("X-Request-ID")
		if id == "" {
			id = uuid.NewString()
		}
		ctx.Locals("requestID", id)
		ctx.Set("X-Request-ID", id)
		return ctx.Next()
	}
}
