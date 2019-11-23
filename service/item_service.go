package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

type ItemService struct {
	itemRepository repository.ItemContractRepository
}

func NewItemService(rs repository.ItemContractRepository) ItemContractService {
	return &ItemService{rs}
}

func (r ItemService) All(itemId int64) ([]model.Item, error) {
	items, err := r.itemRepository.All(itemId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r ItemService) Find(id int64) (model.Item, error) {
	item, err := r.itemRepository.Find(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r ItemService) Store(receiptId int64, item model.Item) (model.Item, error) {
	item, err := r.itemRepository.Store(receiptId, item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r ItemService) Update(item model.Item) (model.Item, error) {
	item, err := r.itemRepository.Update(item)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r ItemService) Destroy(id int64) (model.Item, error) {
	item, err := r.itemRepository.Destroy(id)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}
