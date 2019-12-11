package transformer

import (
	"time"

	"github.com/andrewesteves/tapagguapi/model"
)

// ReceiptDataManyTransformer struct
type ReceiptDataManyTransformer struct {
	Receipts []ReceiptTransformer `json:"receipts"`
	CommonTransformer
}

// ReceiptDataOneTransformer struct
type ReceiptDataOneTransformer struct {
	Receipt ReceiptTransformer `json:"receipt"`
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

// TransformOne receipt specified JSON
func (rf ReceiptTransformer) TransformOne(receipt model.Receipt) ReceiptDataOneTransformer {
	var newReceipts []ReceiptTransformer
	var newData ReceiptDataOneTransformer
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
	if receipt.Company.ID > 0 {
		newReceipt.Company = CompanyTransformer{}.TransformOne(receipt.Company)
	}
	newReceipts = append(newReceipts, newReceipt)
	newData.Receipt = newReceipt
	return newData
}

// TransformMany receipt specified JSON
func (rf ReceiptTransformer) TransformMany(receipts []model.Receipt, values map[string]string) ReceiptDataManyTransformer {
	var newReceipts []ReceiptTransformer
	var newData ReceiptDataManyTransformer
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
		if receipt.Company.ID > 0 {
			newReceipt.Company = CompanyTransformer{}.TransformOne(receipt.Company)
		}
		newReceipts = append(newReceipts, newReceipt)
	}
	if len(newReceipts) > 0 {
		newData.Receipts = newReceipts
	} else {
		newData.Receipts = make([]ReceiptTransformer, 0)
	}
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
