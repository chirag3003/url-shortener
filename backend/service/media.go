package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/chirag3003/go-backend-template/config"
	"github.com/chirag3003/go-backend-template/helpers"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog"
)

// MediaService handles media upload business logic.
type MediaService struct {
	mediaRepo repository.MediaRepository
	s3Repo    repository.S3Repository
	cfg       *config.Config
	log       zerolog.Logger
}

// NewMediaService creates a new MediaService.
func NewMediaService(
	mediaRepo repository.MediaRepository,
	s3Repo repository.S3Repository,
	cfg *config.Config,
	log zerolog.Logger,
) *MediaService {
	return &MediaService{
		mediaRepo: mediaRepo,
		s3Repo:    s3Repo,
		cfg:       cfg,
		log:       log.With().Str("service", "media").Logger(),
	}
}

// UploadFiles processes and uploads multiple files to S3, storing metadata in PostgreSQL.
func (s *MediaService) UploadFiles(ctx context.Context, files []*multipart.FileHeader) ([]string, error) {
	var fileURLs []string

	for _, file := range files {
		name, err := gonanoid.New(10)
		if err != nil {
			s.log.Error().Err(err).Msg("failed to generate nanoid")
			return nil, apperror.Internal(err)
		}

		var reader io.Reader
		reader, err = helpers.OptimiseImage(*file)
		if err != nil {
			s.log.Warn().Err(err).Str("filename", file.Filename).Msg("image optimization failed, using original")
			// Fall back to original file if optimization fails
			f, openErr := file.Open()
			if openErr != nil {
				s.log.Error().Err(openErr).Msg("failed to open file")
				return nil, apperror.Internal(openErr)
			}
			defer f.Close()
			reader = f
		}

		key := fmt.Sprintf("%s/%s%s.webp", s.cfg.S3Folder, name, file.Filename)

		res, err := s.s3Repo.Upload(ctx, key, reader)
		if err != nil {
			s.log.Error().Err(err).Str("key", key).Msg("failed to upload to S3")
			return nil, apperror.Internal(err)
		}

		fileURLs = append(fileURLs, res.Location)

		mediaID, err := idgen.NewID()
		if err != nil {
			s.log.Error().Err(err).Msg("failed to generate media id")
			return nil, apperror.Internal(err)
		}

		if err := s.mediaRepo.CreateMedia(ctx, &models.Media{
			ID:   mediaID,
			Key:  key,
			Url:  res.Location,
			Etag: *res.ETag,
			Mime: file.Header.Get("Content-Type"),
			Size: file.Size,
		}); err != nil {
			s.log.Error().Err(err).Str("key", key).Msg("failed to save media metadata")
			return nil, apperror.Internal(err)
		}
	}

	return fileURLs, nil
}
