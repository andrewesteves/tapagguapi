package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

type ReceiptService struct {
	receiptRepository repository.ReceiptContractRepository
}

func NewReceiptService(rs repository.ReceiptContractRepository) ReceiptContractService {
	return &ReceiptService{rs}
}

func (r ReceiptService) All() ([]model.Receipt, error) {
	return []model.Receipt{}, nil
}

func (r ReceiptService) Find() (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptService) Store(receipt model.Receipt) (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptService) Update(receipt model.Receipt) (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptService) Destroy(id int64) (model.Receipt, error) {
	return model.Receipt{}, nil
}
