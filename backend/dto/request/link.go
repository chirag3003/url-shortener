package request

// CreateLinkRequest represents the short-link create payload.
type CreateLinkRequest struct {
	LongURL      string `json:"longUrl" validate:"required,url,max=2048"`
	CustomAlias  string `json:"customAlias,omitempty" validate:"omitempty,min=3,max=64,alphanum"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
	RedirectType int16  `json:"redirectType,omitempty" validate:"omitempty,oneof=301 302"`
}

// UpdateLinkRequest represents editable short-link fields.
type UpdateLinkRequest struct {
	LongURL      string `json:"longUrl,omitempty" validate:"omitempty,url,max=2048"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
	RedirectType int16  `json:"redirectType,omitempty" validate:"omitempty,oneof=301 302"`
	IsActive     *bool  `json:"isActive,omitempty"`
}

// ListLinksRequest represents paginated list query params.
type ListLinksRequest struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=128"`
}
