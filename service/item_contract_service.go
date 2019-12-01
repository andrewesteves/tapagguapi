package service

import "github.com/andrewesteves/tapagguapi/model"

// ItemContractService contract
type ItemContractService interface {
	All(receiptID int64) ([]model.Item, error)
	Find(id int64) (model.Item, error)
	Store(receiptID int64, r model.Item) (model.Item, error)
	Update(r model.Item) (model.Item, error)
	Destroy(int int64) (model.Item, error)
}
