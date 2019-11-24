package model

import "time"

type Item struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" xml:"prod>xProd"`
	Price     float64   `json:"price" xml:"prod>vUnCom"`
	Quantity  float64   `json:"quantity" xml:"prod>qCom"`
	Total     float64   `json:"total" xml:"prod>vProd"`
	Tax       float64   `json:"tax" xml:"imposto>vTotTrib"`
	Measure   string    `json:"measure" xml:"prod>uCom"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
