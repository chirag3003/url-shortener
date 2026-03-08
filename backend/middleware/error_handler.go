package middleware

import (
	"errors"

	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

// ErrorHandler returns a Fiber error handler that maps AppErrors to JSON responses.
func ErrorHandler(log zerolog.Logger) fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			if appErr.Err != nil {
				log.Error().Err(appErr.Err).
					Str("code", appErr.Code).
					Str("path", ctx.Path()).
					Str("method", ctx.Method()).
					Msg(appErr.Message)
			}
			return ctx.Status(appErr.Status).JSON(response.Err(appErr.Code, appErr.Message))
		}

		// Handle Fiber errors (404, etc.)
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(response.Err("HTTP_ERROR", fiberErr.Message))
		}

		// Unknown error — log and return generic 500
		log.Error().Err(err).
			Str("path", ctx.Path()).
			Str("method", ctx.Method()).
			Msg("unhandled error")
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			response.Err("INTERNAL_ERROR", "Internal server error"),
		)
	}
}
