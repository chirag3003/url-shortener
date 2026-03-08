package models

import "time"

// Link represents a shortened URL mapping.
type Link struct {
	ID           int64      `json:"id,omitempty"`
	UserID       *int64     `json:"userId,omitempty"`
	LongURL      string     `json:"longUrl"`
	ShortCode    string     `json:"shortCode"`
	RedirectType int16      `json:"redirectType"`
	ExpiresAt    *time.Time `json:"expiresAt,omitempty"`
	IsActive     bool       `json:"isActive"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
