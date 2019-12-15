package repository

import "github.com/andrewesteves/tapagguapi/model"

// UserContractRepository contract
type UserContractRepository interface {
	All() ([]model.User, error)
	Find(id int64) (model.User, error)
	Store(u model.User) (model.User, error)
	Update(u model.User) (model.User, error)
	Destroy(int int64) (model.User, error)
	FindBy(field string, value interface{}) (model.User, error)
	FindByArgs(args map[string]interface{}) (model.User, error)
}
