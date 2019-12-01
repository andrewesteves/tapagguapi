package model

import "time"

// Receipt model
type Receipt struct {
	ID        int64     `json:"id"`
	Category  Category  `json:"category"`
	User      User      `json:"user"`
	Company   Company   `json:"company" xml:"proc>nfeProc>NFe>infNFe>emit"`
	Title     string    `json:"title"`
	Tax       float64   `json:"tax" xml:"proc>nfeProc>NFe>infNFe>total>ICMSTot>vTotTrib"`
	Extra     float64   `json:"extra" xml:"proc>nfeProc>NFe>infNFe>total>ICMSTot>vOutro"`
	Discount  float64   `json:"discount" xml:"proc>nfeProc>NFe>infNFe>total>ICMSTot>vDesc"`
	Total     float64   `json:"total" xml:"proc>nfeProc>NFe>infNFe>total>ICMSTot>vNF"`
	Items     []Item    `json:"items" xml:"proc>nfeProc>NFe>infNFe>det"`
	URL       string    `json:"url" xml:"proc>nfeProc>NFe>infNFeSupl>qrCode"`
	AccessKey string    `json:"accessKey" xml:"proc>nfeProc>protNFe>infProt>chNFe"`
	IssuedAt  time.Time `json:"issuedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
