package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

// RequestLogger logs each request using zerolog.
func RequestLogger(log zerolog.Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		duration := time.Since(start)
		status := ctx.Response().StatusCode()

		event := log.Info()
		if status >= 400 {
			event = log.Warn()
		}
		if status >= 500 {
			event = log.Error()
		}

		requestID, _ := ctx.Locals("requestID").(string)

		event.
			Str("method", ctx.Method()).
			Str("path", ctx.Path()).
			Int("status", status).
			Dur("duration", duration).
			Str("ip", ctx.IP()).
			Str("request_id", requestID).
			Msg("request")

		return err
	}
}
