package model

import "time"

type Item struct {
	Title     string  `json:"title" xml:"prod>xProd"`
	Price     float64 `json:"price" xml:"prod>vUnCom"`
	Qty       float64 `json:"quantity" xml:"prod>qCom"`
	Total     float64 `json:"total" xml:"prod>vProd"`
	Tax       float64 `json:"tax" xml:"imposto>vTotTrib"`
	Measure   string  `json:"measure" xml:"prod>uCom"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
