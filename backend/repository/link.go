package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination=mock/mock_link.go -package=mock github.com/chirag3003/go-backend-template/repository LinkRepository

// LinkRepository defines short-link persistence operations.
type LinkRepository interface {
	CreateLink(ctx context.Context, link *models.Link) error
	GetLinkByID(ctx context.Context, id int64) (*models.Link, error)
	GetLinkByCode(ctx context.Context, code string) (*models.Link, error)
	ListLinksByUser(ctx context.Context, userID int64, page int, limit int, search string) ([]*models.Link, int64, error)
	UpdateLink(ctx context.Context, link *models.Link) error
	DeleteLink(ctx context.Context, id int64, userID int64) error
	AliasExists(ctx context.Context, alias string) (bool, error)
}

type linkRepository struct {
	conn db.Connection
}

// NewLinkRepository creates a new LinkRepository.
func NewLinkRepository(conn db.Connection) LinkRepository {
	return &linkRepository{conn: conn}
}

func (r *linkRepository) CreateLink(ctx context.Context, link *models.Link) error {
	const q = `
		INSERT INTO links (
			id, user_id, long_url, short_code, redirect_type, expires_at, is_active, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`
	_, err := r.conn.Pool().Exec(
		ctx,
		q,
		link.ID,
		link.UserID,
		link.LongURL,
		link.ShortCode,
		link.RedirectType,
		link.ExpiresAt,
		link.IsActive,
	)
	return err
}

func (r *linkRepository) GetLinkByID(ctx context.Context, id int64) (*models.Link, error) {
	const q = `
		SELECT id, user_id, long_url, short_code, redirect_type, expires_at, is_active, created_at, updated_at
		FROM links
		WHERE id = $1`

	var link models.Link
	var userID *int64
	err := r.conn.Pool().QueryRow(ctx, q, id).Scan(
		&link.ID,
		&userID,
		&link.LongURL,
		&link.ShortCode,
		&link.RedirectType,
		&link.ExpiresAt,
		&link.IsActive,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	link.UserID = userID
	return &link, nil
}

func (r *linkRepository) GetLinkByCode(ctx context.Context, code string) (*models.Link, error) {
	const q = `
		SELECT id, user_id, long_url, short_code, redirect_type, expires_at, is_active, created_at, updated_at
		FROM links
		WHERE short_code = $1`

	var link models.Link
	var userID *int64
	err := r.conn.Pool().QueryRow(ctx, q, code).Scan(
		&link.ID,
		&userID,
		&link.LongURL,
		&link.ShortCode,
		&link.RedirectType,
		&link.ExpiresAt,
		&link.IsActive,
		&link.CreatedAt,
		&link.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	link.UserID = userID
	return &link, nil
}

func (r *linkRepository) ListLinksByUser(ctx context.Context, userID int64, page int, limit int, search string) ([]*models.Link, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var where []string
	args := []interface{}{userID}
	where = append(where, fmt.Sprintf("user_id = $%d", len(args)))

	if search != "" {
		args = append(args, "%"+strings.ToLower(search)+"%")
		where = append(where, fmt.Sprintf("(LOWER(short_code) LIKE $%d OR LOWER(long_url) LIKE $%d)", len(args), len(args)))
	}

	whereClause := strings.Join(where, " AND ")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM links WHERE %s", whereClause)

	var total int64
	if err := r.conn.Pool().QueryRow(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	args = append(args, limit, offset)
	listQ := fmt.Sprintf(`
		SELECT id, user_id, long_url, short_code, redirect_type, expires_at, is_active, created_at, updated_at
		FROM links
		WHERE %s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d`, whereClause, len(args)-1, len(args))

	rows, err := r.conn.Pool().Query(ctx, listQ, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]*models.Link, 0)
	for rows.Next() {
		var item models.Link
		var uid *int64
		if err := rows.Scan(
			&item.ID,
			&uid,
			&item.LongURL,
			&item.ShortCode,
			&item.RedirectType,
			&item.ExpiresAt,
			&item.IsActive,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		item.UserID = uid
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *linkRepository) UpdateLink(ctx context.Context, link *models.Link) error {
	const q = `
		UPDATE links
		SET long_url = $2, redirect_type = $3, expires_at = $4, is_active = $5, updated_at = NOW()
		WHERE id = $1`
	_, err := r.conn.Pool().Exec(ctx, q, link.ID, link.LongURL, link.RedirectType, link.ExpiresAt, link.IsActive)
	return err
}

func (r *linkRepository) DeleteLink(ctx context.Context, id int64, userID int64) error {
	const q = `DELETE FROM links WHERE id = $1 AND user_id = $2`
	_, err := r.conn.Pool().Exec(ctx, q, id, userID)
	return err
}

func (r *linkRepository) AliasExists(ctx context.Context, alias string) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM links WHERE short_code = $1)`
	var exists bool
	err := r.conn.Pool().QueryRow(ctx, q, alias).Scan(&exists)
	return exists, err
}

//go:generate mockgen -destination=mock/mock_click.go -package=mock github.com/chirag3003/go-backend-template/repository ClickRepository

// ClickRepository defines click event persistence operations.
type ClickRepository interface {
	CreateClick(ctx context.Context, click *models.ClickEvent) error
	GetSummary(ctx context.Context, linkID int64) (total int64, unique int64, last24h int64, last7d int64, err error)
	GetPreviousSummary(ctx context.Context, linkID int64) (prevTotal int64, prevUnique int64, prev24h int64, prev7d int64, err error)
	GetTimeSeries(ctx context.Context, linkID int64, window string) ([]TimeSeriesPoint, error)
	GetTopBreakdown(ctx context.Context, linkID int64, field string, limit int) ([]BreakdownRow, error)
}

// TimeSeriesPoint represents one analytics bucket.
type TimeSeriesPoint struct {
	Bucket time.Time
	Count  int64
}

// BreakdownRow represents key-count analytics row.
type BreakdownRow struct {
	Key   string
	Count int64
}

type clickRepository struct {
	conn db.Connection
}

// NewClickRepository creates a new ClickRepository.
func NewClickRepository(conn db.Connection) ClickRepository {
	return &clickRepository{conn: conn}
}

func (r *clickRepository) CreateClick(ctx context.Context, click *models.ClickEvent) error {
	const q = `
		INSERT INTO click_events (
			id, link_id, clicked_at, ip_address, user_agent, referrer, country, device_type, browser
		)
		VALUES ($1, $2, NOW(), $3, $4, $5, $6, $7, $8)`
	_, err := r.conn.Pool().Exec(
		ctx,
		q,
		click.ID,
		click.LinkID,
		nullableString(click.IPAddress),
		nullableString(click.UserAgent),
		nullableString(click.Referrer),
		nullableString(click.Country),
		nullableString(click.DeviceType),
		nullableString(click.Browser),
	)
	return err
}

func (r *clickRepository) GetSummary(ctx context.Context, linkID int64) (int64, int64, int64, int64, error) {
	const q = `
		SELECT
			COUNT(*) AS total,
			COUNT(DISTINCT COALESCE(ip_address, '')) AS unique_visitors,
			COUNT(*) FILTER (WHERE clicked_at >= NOW() - INTERVAL '24 hours') AS clicks_last_24h,
			COUNT(*) FILTER (WHERE clicked_at >= NOW() - INTERVAL '7 days') AS clicks_last_7d
		FROM click_events
		WHERE link_id = $1`

	var total int64
	var unique int64
	var last24h int64
	var last7d int64
	err := r.conn.Pool().QueryRow(ctx, q, linkID).Scan(&total, &unique, &last24h, &last7d)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return total, unique, last24h, last7d, nil
}

func (r *clickRepository) GetPreviousSummary(ctx context.Context, linkID int64) (int64, int64, int64, int64, error) {
	const q = `
		SELECT
			COUNT(*) FILTER (WHERE clicked_at < NOW() - INTERVAL '24 hours') AS prev_total,
			COUNT(DISTINCT COALESCE(ip_address, '')) FILTER (WHERE clicked_at < NOW() - INTERVAL '24 hours') AS prev_unique,
			COUNT(*) FILTER (WHERE clicked_at >= NOW() - INTERVAL '48 hours' AND clicked_at < NOW() - INTERVAL '24 hours') AS prev_24h,
			COUNT(*) FILTER (WHERE clicked_at >= NOW() - INTERVAL '14 days' AND clicked_at < NOW() - INTERVAL '7 days') AS prev_7d
		FROM click_events
		WHERE link_id = $1`

	var prevTotal int64
	var prevUnique int64
	var prev24h int64
	var prev7d int64
	err := r.conn.Pool().QueryRow(ctx, q, linkID).Scan(&prevTotal, &prevUnique, &prev24h, &prev7d)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return prevTotal, prevUnique, prev24h, prev7d, nil
}

func (r *clickRepository) GetTimeSeries(ctx context.Context, linkID int64, window string) ([]TimeSeriesPoint, error) {
	bucket := "day"
	interval := "30 days"
	if window == "24h" {
		bucket = "hour"
		interval = "24 hours"
	}
	if window == "7d" {
		bucket = "day"
		interval = "7 days"
	}

	q := fmt.Sprintf(`
		SELECT DATE_TRUNC('%s', clicked_at) AS bucket, COUNT(*) AS clicks
		FROM click_events
		WHERE link_id = $1 AND clicked_at >= NOW() - INTERVAL '%s'
		GROUP BY bucket
		ORDER BY bucket ASC`, bucket, interval)

	rows, err := r.conn.Pool().Query(ctx, q, linkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := make([]TimeSeriesPoint, 0)
	for rows.Next() {
		var p TimeSeriesPoint
		if err := rows.Scan(&p.Bucket, &p.Count); err != nil {
			return nil, err
		}
		points = append(points, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func (r *clickRepository) GetTopBreakdown(ctx context.Context, linkID int64, field string, limit int) ([]BreakdownRow, error) {
	allowed := map[string]bool{
		"referrer":    true,
		"device_type": true,
		"browser":     true,
		"country":     true,
	}
	if !allowed[field] {
		return nil, fmt.Errorf("unsupported breakdown field: %s", field)
	}
	if limit < 1 {
		limit = 10
	}

	q := fmt.Sprintf(`
		SELECT COALESCE(NULLIF(%s, ''), 'unknown') AS key, COUNT(*) AS count
		FROM click_events
		WHERE link_id = $1
		GROUP BY key
		ORDER BY count DESC
		LIMIT $2`, field)

	rows, err := r.conn.Pool().Query(ctx, q, linkID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]BreakdownRow, 0)
	for rows.Next() {
		var row BreakdownRow
		if err := rows.Scan(&row.Key, &row.Count); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
