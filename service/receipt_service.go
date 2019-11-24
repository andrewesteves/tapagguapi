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
	receipts, err := r.receiptRepository.All()
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

func (r ReceiptService) Find(id int64) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Find(id)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

func (r ReceiptService) Store(receipt model.Receipt) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Store(receipt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

func (r ReceiptService) Update(receipt model.Receipt) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Update(receipt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

func (r ReceiptService) Destroy(id int64) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Destroy(id)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}
