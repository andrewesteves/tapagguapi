package model

import "time"

type Receipt struct {
	ID        int64
	Title     string
	Tax       float64
	Total     float64
	Items     []Item
	URL       string
	IssueAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
