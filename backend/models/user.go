package models

import "time"

// User represents the user record in PostgreSQL.
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PhoneNo   string    `json:"phoneNo,omitempty"`
	Hash      string    `json:"-"`
	AvatarURL string    `json:"avatarUrl,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
