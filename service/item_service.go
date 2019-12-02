package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

// ItemService struct
type ItemService struct {
	itemRepository repository.ItemContractRepository
}

// NewItemService new item service
func NewItemService(rs repository.ItemContractRepository) ItemContractService {
	return &ItemService{rs}
}

// All items service
func (r ItemService) All(receiptID int64) ([]model.Item, error) {
	items, err := r.itemRepository.All(receiptID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Find item service
func (r ItemService) Find(id int64) (model.Item, error) {
	item, err := r.itemRepository.Find(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// Store item service
func (r ItemService) Store(receiptID int64, item model.Item) (model.Item, error) {
	item, err := r.itemRepository.Store(receiptID, item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// Update item service
func (r ItemService) Update(item model.Item) (model.Item, error) {
	item, err := r.itemRepository.Update(item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// Destroy item service
func (r ItemService) Destroy(id int64) (model.Item, error) {
	item, err := r.itemRepository.Destroy(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}
