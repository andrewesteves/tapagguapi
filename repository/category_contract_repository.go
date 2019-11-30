package repository

import "github.com/andrewesteves/tapagguapi/model"

// CategoryContractRepository contract
type CategoryContractRepository interface {
	All(user model.User) ([]model.Category, error)
	Find(id int64) (model.Category, error)
	Store(r model.Category) (model.Category, error)
	Update(r model.Category) (model.Category, error)
	Destroy(int int64) (model.Category, error)
}
