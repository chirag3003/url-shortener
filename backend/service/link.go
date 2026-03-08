package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/cache"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/pkg/validate"
	"github.com/chirag3003/go-backend-template/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rs/zerolog"
)

// LinkService handles short-link business logic.
type LinkService struct {
	linkRepo repository.LinkRepository
	baseURL  string
	log      zerolog.Logger
}

// NewLinkService creates a new LinkService.
func NewLinkService(linkRepo repository.LinkRepository, baseURL string, log zerolog.Logger) *LinkService {
	trimmedBase := strings.TrimSuffix(strings.TrimSpace(baseURL), "/")
	if trimmedBase == "" {
		trimmedBase = "http://localhost:5000"
	}

	return &LinkService{
		linkRepo: linkRepo,
		baseURL:  trimmedBase,
		log:      log.With().Str("service", "link").Logger(),
	}
}

// Create creates a short link.
func (s *LinkService) Create(ctx context.Context, userID *int64, req *request.CreateLinkRequest) (*response.LinkResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, apperror.ValidationError(err.Error())
	}

	shortCode := req.CustomAlias
	if shortCode == "" {
		generated, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 8)
		if err != nil {
			s.log.Error().Err(err).Msg("failed to generate short code")
			return nil, apperror.Internal(err)
		}
		shortCode = generated
	}

	exists, err := s.linkRepo.AliasExists(ctx, shortCode)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if exists {
		return nil, apperror.New("ALIAS_TAKEN", "custom alias is already taken", 409)
	}

	linkID, err := idgen.NewID()
	if err != nil {
		return nil, apperror.Internal(err)
	}

	redirectType := req.RedirectType
	if redirectType == 0 {
		redirectType = 302
	}

	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		t, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return nil, apperror.BadRequest("expiresAt must be a valid RFC3339 timestamp")
		}
		expiresAt = &t
	}

	link := &models.Link{
		ID:           linkID,
		UserID:       userID,
		LongURL:      req.LongURL,
		ShortCode:    shortCode,
		RedirectType: redirectType,
		ExpiresAt:    expiresAt,
		IsActive:     true,
	}

	if err := s.linkRepo.CreateLink(ctx, link); err != nil {
		s.log.Error().Err(err).Msg("failed to create link")
		return nil, apperror.Internal(err)
	}

	stored, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if stored == nil {
		return nil, apperror.Internal(apperror.ErrInternal)
	}

	return s.toLinkResponse(stored), nil
}

// GetByID retrieves a short link by ID and validates ownership for authenticated users.
func (s *LinkService) GetByID(ctx context.Context, userID int64, linkID int64) (*response.LinkResponse, error) {
	link, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if link == nil {
		return nil, apperror.NotFound("link not found")
	}
	if link.UserID == nil || *link.UserID != userID {
		return nil, apperror.ErrForbidden
	}
	return s.toLinkResponse(link), nil
}

// ListByUser returns paginated links owned by the user.
func (s *LinkService) ListByUser(ctx context.Context, userID int64, req *request.ListLinksRequest) (*response.PaginatedLinksResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, apperror.ValidationError(err.Error())
	}

	page := req.Page
	if page == 0 {
		page = 1
	}
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}

	links, total, err := s.linkRepo.ListLinksByUser(ctx, userID, page, limit, req.Search)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	items := make([]response.LinkResponse, 0, len(links))
	for _, item := range links {
		items = append(items, *s.toLinkResponse(item))
	}

	return &response.PaginatedLinksResponse{
		Items: items,
		Page:  page,
		Limit: limit,
		Total: total,
	}, nil
}

// Update updates editable fields for a link owned by the user.
func (s *LinkService) Update(ctx context.Context, userID int64, linkID int64, req *request.UpdateLinkRequest) (*response.LinkResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, apperror.ValidationError(err.Error())
	}

	link, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if link == nil {
		return nil, apperror.NotFound("link not found")
	}
	if link.UserID == nil || *link.UserID != userID {
		return nil, apperror.ErrForbidden
	}

	if req.LongURL != "" {
		link.LongURL = req.LongURL
	}
	if req.RedirectType != 0 {
		link.RedirectType = req.RedirectType
	}
	if req.ExpiresAt != "" {
		t, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return nil, apperror.BadRequest("expiresAt must be a valid RFC3339 timestamp")
		}
		link.ExpiresAt = &t
	}
	if req.IsActive != nil {
		link.IsActive = *req.IsActive
	}

	if err := s.linkRepo.UpdateLink(ctx, link); err != nil {
		return nil, apperror.Internal(err)
	}

	// Invalidate cache
	_ = cache.Delete(ctx, "link:"+link.ShortCode)

	updated, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if updated == nil {
		return nil, apperror.NotFound("link not found")
	}

	return s.toLinkResponse(updated), nil
}

// Delete removes one link owned by the user.
func (s *LinkService) Delete(ctx context.Context, userID int64, linkID int64) error {
	link, err := s.linkRepo.GetLinkByID(ctx, linkID)
	if err != nil {
		return apperror.Internal(err)
	}
	if link == nil {
		return apperror.NotFound("link not found")
	}
	if link.UserID == nil || *link.UserID != userID {
		return apperror.ErrForbidden
	}

	if err := s.linkRepo.DeleteLink(ctx, linkID, userID); err != nil {
		return apperror.Internal(err)
	}

	// Invalidate cache
	_ = cache.Delete(ctx, "link:"+link.ShortCode)

	return nil
}

// CheckAliasAvailability checks whether a custom alias is available.
func (s *LinkService) CheckAliasAvailability(ctx context.Context, alias string) (*response.AliasAvailabilityResponse, error) {
	alias = strings.TrimSpace(alias)
	if alias == "" {
		return nil, apperror.BadRequest("alias is required")
	}

	exists, err := s.linkRepo.AliasExists(ctx, alias)
	if err != nil {
		return nil, apperror.Internal(err)
	}

	return &response.AliasAvailabilityResponse{
		Alias:     alias,
		Available: !exists,
	}, nil
}

// ResolveForRedirect resolves an active and non-expired short code.
func (s *LinkService) ResolveForRedirect(ctx context.Context, code string) (*models.Link, error) {
	// 1. Try Cache
	cached, err := cache.Get(ctx, "link:"+code)
	if err == nil && cached != "" {
		// For now, we store just the URL in cache for simplicity in this prototype.
		// In a full implementation, we'd store a JSON of models.Link to handle RedirectType, etc.
		return &models.Link{
			LongURL:      cached,
			ShortCode:    code,
			RedirectType: 302,
			IsActive:     true,
		}, nil
	}

	// 2. Try DB
	link, err := s.linkRepo.GetLinkByCode(ctx, code)
	if err != nil {
		return nil, apperror.Internal(err)
	}
	if link == nil || !link.IsActive {
		return nil, apperror.NotFound("short link not found")
	}
	if link.ExpiresAt != nil && time.Now().After(*link.ExpiresAt) {
		return nil, apperror.NotFound("short link expired")
	}

	// 3. Populate Cache (TTL 1 hour)
	_ = cache.Set(ctx, "link:"+code, link.LongURL, 1*time.Hour)

	return link, nil
}

func (s *LinkService) toLinkResponse(link *models.Link) *response.LinkResponse {
	resp := &response.LinkResponse{
		ID:           strconv.FormatInt(link.ID, 10),
		LongURL:      link.LongURL,
		ShortCode:    link.ShortCode,
		ShortURL:     s.baseURL + "/" + link.ShortCode,
		RedirectType: link.RedirectType,
		IsActive:     link.IsActive,
		CreatedAt:    link.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    link.UpdatedAt.Format(time.RFC3339),
	}
	if link.UserID != nil {
		resp.UserID = strconv.FormatInt(*link.UserID, 10)
	}
	if link.ExpiresAt != nil {
		resp.ExpiresAt = link.ExpiresAt.Format(time.RFC3339)
	}
	return resp
}
