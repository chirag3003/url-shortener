package repository

import (
	"context"
	"errors"

	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination=mock/mock_media.go -package=mock github.com/chirag3003/go-backend-template/repository MediaRepository

// MediaRepository defines the interface for media data access.
type MediaRepository interface {
	CreateMedia(ctx context.Context, media *models.Media) error
	GetMediaByID(ctx context.Context, id int64) (*models.Media, error)
	GetMediaByKey(ctx context.Context, key string) (*models.Media, error)
	UpdateMedia(ctx context.Context, media *models.Media) error
}

type mediaRepository struct {
	conn db.Connection
}

// NewMediaRepository creates a new MediaRepository.
func NewMediaRepository(conn db.Connection) MediaRepository {
	return &mediaRepository{conn: conn}
}

func (r *mediaRepository) CreateMedia(ctx context.Context, media *models.Media) error {
	const q = `
		INSERT INTO media (id, key, etag, size, mime, url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`
	_, err := r.conn.Pool().Exec(ctx, q, media.ID, media.Key, media.Etag, media.Size, media.Mime, media.Url)
	return err
}

func (r *mediaRepository) GetMediaByID(ctx context.Context, id int64) (*models.Media, error) {
	const q = `
		SELECT id, key, etag, size, mime, url, created_at, updated_at
		FROM media
		WHERE id = $1`

	var media models.Media
	err := r.conn.Pool().QueryRow(ctx, q, id).Scan(
		&media.ID,
		&media.Key,
		&media.Etag,
		&media.Size,
		&media.Mime,
		&media.Url,
		&media.CreatedAt,
		&media.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepository) GetMediaByKey(ctx context.Context, key string) (*models.Media, error) {
	const q = `
		SELECT id, key, etag, size, mime, url, created_at, updated_at
		FROM media
		WHERE key = $1`

	var media models.Media
	err := r.conn.Pool().QueryRow(ctx, q, key).Scan(
		&media.ID,
		&media.Key,
		&media.Etag,
		&media.Size,
		&media.Mime,
		&media.Url,
		&media.CreatedAt,
		&media.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func (r *mediaRepository) UpdateMedia(ctx context.Context, media *models.Media) error {
	const q = `
		UPDATE media
		SET key = $2, etag = $3, size = $4, mime = $5, url = $6, updated_at = NOW()
		WHERE id = $1`
	_, err := r.conn.Pool().Exec(ctx, q, media.ID, media.Key, media.Etag, media.Size, media.Mime, media.Url)
	return err
}
