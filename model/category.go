package model

import "time"

// Category model
type Category struct {
	ID        int64     `json:"id,omitempty"`
	User      User      `json:"user"`
	Title     string    `json:"title,omitempty"`
	Icon      string    `json:"icon,omitempty"`
	Total     float64   `json:"total,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
