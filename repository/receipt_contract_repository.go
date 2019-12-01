package repository

import "github.com/andrewesteves/tapagguapi/model"

// ReceiptContractRepository contract
type ReceiptContractRepository interface {
	All(user model.User) ([]model.Receipt, error)
	Find(id int64) (model.Receipt, error)
	Store(r model.Receipt) (model.Receipt, error)
	Update(r model.Receipt) (model.Receipt, error)
	Destroy(int int64) (model.Receipt, error)
	FindManyBy(field string, value interface{}) ([]model.Receipt, error)
	RetrieveStore(r model.Receipt) (model.Receipt, error)
	Count(user model.User) (int, error)
}
