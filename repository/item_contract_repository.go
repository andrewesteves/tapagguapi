package repository

import "github.com/andrewesteves/tapagguapi/model"

// ItemContractRepository contract
type ItemContractRepository interface {
	All(receiptID int64) ([]model.Item, error)
	Find(id int64) (model.Item, error)
	Store(receiptID int64, r model.Item) (model.Item, error)
	Update(r model.Item) (model.Item, error)
	Destroy(int int64) (model.Item, error)
}
