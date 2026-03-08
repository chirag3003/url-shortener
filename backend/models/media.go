package models

import "time"

// Media represents the media record in PostgreSQL.
type Media struct {
	ID        int64     `json:"id,omitempty"`
	Key       string    `json:"key"`
	Etag      string    `json:"etag"`
	Size      int64     `json:"size"`
	Mime      string    `json:"mime"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
