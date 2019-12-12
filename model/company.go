package model

import "time"

// Company model
type Company struct {
	ID        int64     `json:"id,omitempty"`
	User      User      `json:"-"`
	CNPJ      string    `json:"cnpj,omitempty" xml:"CNPJ"`
	Name      string    `json:"name,omitempty" xml:"xNome"`
	Title     string    `json:"title,omitempty" xml:"xFant"`
	Street    string    `json:"street,omitempty" xml:"enderEmit>xLgr"`
	Number    string    `json:"number,omitempty" xml:"enderEmit>nro"`
	District  string    `json:"district,omitempty" xml:"enderEmit>xBairro"`
	City      string    `json:"city,omitempty" xml:"enderEmit>xMun"`
	State     string    `json:"state,omitempty" xml:"enderEmit>UF"`
	Zipcode   string    `json:"zipcode,omitempty" xml:"enderEmit>CEP"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
