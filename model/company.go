package model

import "time"

// Company model
type Company struct {
	ID        int64     `json:"id,omitempty"`
	User      User      `json:"-"`
	CNPJ      string    `json:"cnpj,omitempty" xml:"CNPJ"`
	Name      string    `json:"name,omitempty" xml:"xNome"`
	Title     string    `json:"title,omitempty" xml:"xFant"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
