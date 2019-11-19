package model

import "time"

type Item struct {
	Receipt   Receipt
	Title     string
	Price     float64
	Qty       float64
	Total     float64
	Tax       float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
