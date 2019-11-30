package model

type Company struct {
	ID    int64  `json:"id,omitempty"`
	CNPJ  string `json:"cnpj,omitempty" xml:"CNPJ"`
	Name  string `json:"name,omitempty" xml:"xNome"`
	Title string `json:"title,omitempty" xml:"xFant"`
}
