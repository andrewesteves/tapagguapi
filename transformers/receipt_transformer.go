package transformers

import (
	"time"

	"github.com/andrewesteves/tapagguapi/model"
)

// ReceiptDataTransformer struct
type ReceiptDataTransformer struct {
	Receips []ReceiptTransformer `json:"receipts"`
	CommonTransformer
}

// ReceiptTransformer struct
type ReceiptTransformer struct {
	ID        int64        `json:"id"`
	Category  interface{}  `json:"category,omitempty"`
	Company   interface{}  `json:"company,omitempty"`
	Title     string       `json:"title"`
	Tax       float64      `json:"tax"`
	Extra     float64      `json:"extra"`
	Discount  float64      `json:"discount"`
	Total     float64      `json:"total"`
	Items     []model.Item `json:"items"`
	URL       string       `json:"url"`
	AccessKey string       `json:"accessKey"`
	IssuedAt  time.Time    `json:"issuedAt"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

// TransformMany receipt specified JSON
func (rf ReceiptTransformer) TransformMany(receipts []model.Receipt, values map[string]string) ReceiptDataTransformer {
	var newReceipts []ReceiptTransformer
	var newData ReceiptDataTransformer
	for _, receipt := range receipts {
		var newReceipt ReceiptTransformer
		newReceipt.ID = receipt.ID
		newReceipt.Company = receipt.Company
		newReceipt.Title = receipt.Title
		newReceipt.Tax = receipt.Tax
		newReceipt.Extra = receipt.Extra
		newReceipt.Discount = receipt.Discount
		newReceipt.Total = receipt.Total
		newReceipt.Items = receipt.Items
		newReceipt.URL = receipt.URL
		newReceipt.AccessKey = receipt.AccessKey
		newReceipt.IssuedAt = receipt.IssuedAt
		newReceipt.CreatedAt = receipt.CreatedAt
		newReceipt.UpdatedAt = receipt.UpdatedAt
		if receipt.Category.ID > 0 {
			newReceipt.Category = CategoryTransformer{}.TransformOne(receipt.Category)
		}
		newReceipts = append(newReceipts, newReceipt)
	}
	newData.Receips = newReceipts

	if current, ok := values["current"]; ok {
		newData.Current = current
	}
	if prev, ok := values["prev"]; ok {
		newData.Prev = prev
	}
	if next, ok := values["next"]; ok {
		newData.Next = next
	}
	if total, ok := values["total"]; ok {
		newData.Total = total
	}
	return newData
}
