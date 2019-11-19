package service

import "github.com/andrewesteves/tapagguapi/model"

type ReceiptContractService interface {
	All() ([]model.Receipt, error)
	Find() (model.Receipt, error)
	Store(r model.Receipt) (model.Receipt, error)
	Update(r model.Receipt) (model.Receipt, error)
	Destroy(int int64) (model.Receipt, error)
}
