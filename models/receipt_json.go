package models

type ReceiptJSON struct {
	Ide       IdentificationJSON `json:"identification"`
	Iss       IssuerJSON         `json:"issuer"`
	Items     []ItemJSON         `json:"items"`
	Summary   TotalJSON          `json:"summary"`
	ConsultAt string             `json:"consult_at"`
}

type IdentificationJSON struct {
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

type IssuerJSON struct {
	CNPJ    string           `json:"cnpj"`
	Name    string           `json:"name"`
	Title   string           `json:"title"`
	Address IsserAddressJSON `json:"address"`
}

type IsserAddressJSON struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	District string `json:"district"`
	City     string `json:"city"`
	UF       string `json:"uf"`
	Zipcode  string `json:"zipcode"`
}

type ItemJSON struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Qty   float64 `json:"quantity"`
	Total float64 `json:"total"`
	Tax   float64 `json:"tax"`
}

type TotalJSON struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}
