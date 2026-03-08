package models

import "time"

// ClickEvent represents one redirect analytics event.
type ClickEvent struct {
	ID         int64     `json:"id,omitempty"`
	LinkID     int64     `json:"linkId"`
	ClickedAt  time.Time `json:"clickedAt"`
	IPAddress  string    `json:"ipAddress,omitempty"`
	UserAgent  string    `json:"userAgent,omitempty"`
	Referrer   string    `json:"referrer,omitempty"`
	Country    string    `json:"country,omitempty"`
	DeviceType string    `json:"deviceType,omitempty"`
	Browser    string    `json:"browser,omitempty"`
}
