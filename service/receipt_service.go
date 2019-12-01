package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

// ReceiptService struct
type ReceiptService struct {
	receiptRepository repository.ReceiptContractRepository
}

// NewReceiptService new receipt service
func NewReceiptService(rs repository.ReceiptContractRepository) ReceiptContractService {
	return &ReceiptService{rs}
}

// All receipts service
func (r ReceiptService) All(user model.User, values map[string]string) ([]model.Receipt, error) {
	receipts, err := r.receiptRepository.All(user, values)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

// Find receipt service
func (r ReceiptService) Find(id int64) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Find(id)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

// Store receipt service
func (r ReceiptService) Store(receipt model.Receipt) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Store(receipt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

// Update receipt service
func (r ReceiptService) Update(receipt model.Receipt) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Update(receipt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

// Destroy receipt service
func (r ReceiptService) Destroy(id int64) (model.Receipt, error) {
	receipt, err := r.receiptRepository.Destroy(id)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

// FindManyBy receipt service
func (r ReceiptService) FindManyBy(field string, value interface{}) ([]model.Receipt, error) {
	receipts, err := r.receiptRepository.FindManyBy(field, value)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

// RetrieveStore receipt service
func (r ReceiptService) RetrieveStore(receipt model.Receipt) (model.Receipt, error) {
	receipt, err := r.receiptRepository.RetrieveStore(receipt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}
