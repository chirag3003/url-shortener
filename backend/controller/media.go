package controller

import (
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/service"
	"github.com/gofiber/fiber/v3"
)

// MediaController handles media HTTP requests.
type MediaController struct {
	mediaService *service.MediaService
}

// NewMediaController creates a new MediaController.
func NewMediaController(mediaService *service.MediaService) *MediaController {
	return &MediaController{mediaService: mediaService}
}

// Upload handles POST /api/v1/media/upload
func (c *MediaController) Upload(ctx fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return apperror.BadRequest("invalid multipart form")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return apperror.BadRequest("no files provided")
	}

	fileURLs, err := c.mediaService.UploadFiles(ctx.Context(), files)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.OK(
		response.MediaUploadResponse{
			Message: "upload successful",
			Files:   fileURLs,
		},
	))
}
