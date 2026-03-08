package response

// LinkResponse represents short-link data returned via API.
type LinkResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"userId,omitempty"`
	LongURL      string `json:"longUrl"`
	ShortCode    string `json:"shortCode"`
	ShortURL     string `json:"shortUrl"`
	RedirectType int16  `json:"redirectType"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
	IsActive     bool   `json:"isActive"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

// PaginatedLinksResponse represents list links payload.
type PaginatedLinksResponse struct {
	Items []LinkResponse `json:"items"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
	Total int64          `json:"total"`
}

// AliasAvailabilityResponse represents custom alias availability result.
type AliasAvailabilityResponse struct {
	Alias     string `json:"alias"`
	Available bool   `json:"available"`
}
