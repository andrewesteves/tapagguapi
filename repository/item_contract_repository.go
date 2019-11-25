package repository

import "github.com/andrewesteves/tapagguapi/model"

type ItemContractRepository interface {
	Create()
	All(receiptId int64) ([]model.Item, error)
	Find(id int64) (model.Item, error)
	Store(receiptId int64, r model.Item) (model.Item, error)
	Update(r model.Item) (model.Item, error)
	Destroy(int int64) (model.Item, error)
}
