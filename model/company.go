package model

type Company struct {
	ID    int64  `json:"id"`
	CNPJ  string `json:"cnpj" xml:"CNPJ"`
	Name  string `json:"name" xml:"xNome"`
	Title string `json:"title" xml:"xFant"`
}
