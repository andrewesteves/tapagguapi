package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

// CategoryService struct
type CategoryService struct {
	categoryRepository repository.CategoryContractRepository
}

// NewCategoryService new category service
func NewCategoryService(rs repository.CategoryContractRepository) CategoryContractService {
	return &CategoryService{rs}
}

// All categories service
func (r CategoryService) All(user model.User) ([]model.Category, error) {
	categories, err := r.categoryRepository.All(user)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// Find category service
func (r CategoryService) Find(id int64) (model.Category, error) {
	category, err := r.categoryRepository.Find(id)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Store category service
func (r CategoryService) Store(category model.Category) (model.Category, error) {
	category, err := r.categoryRepository.Store(category)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Update category service
func (r CategoryService) Update(category model.Category) (model.Category, error) {
	category, err := r.categoryRepository.Update(category)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Destroy category service
func (r CategoryService) Destroy(id int64) (model.Category, error) {
	category, err := r.categoryRepository.Destroy(id)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}
