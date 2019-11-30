package service

import "github.com/andrewesteves/tapagguapi/model"

// ReceiptContractService contract
type ReceiptContractService interface {
	All(user model.User) ([]model.Receipt, error)
	Find(id int64) (model.Receipt, error)
	Store(r model.Receipt) (model.Receipt, error)
	Update(r model.Receipt) (model.Receipt, error)
	Destroy(int int64) (model.Receipt, error)
}
