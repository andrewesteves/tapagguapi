package service

import "github.com/andrewesteves/tapagguapi/model"

// CategoryContractService contract
type CategoryContractService interface {
	All(user model.User) ([]model.Category, error)
	Find(id int64) (model.Category, error)
	Store(r model.Category) (model.Category, error)
	Update(r model.Category) (model.Category, error)
	Destroy(int int64) (model.Category, error)
}
