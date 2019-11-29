package service

import "github.com/andrewesteves/tapagguapi/model"

type UserContractService interface {
	All() ([]model.User, error)
	Find(id int64) (model.User, error)
	Store(u model.User) (model.User, error)
	Update(u model.User) (model.User, error)
	Destroy(int int64) (model.User, error)
	Login(u model.User) (model.User, error)
	Logout(u model.User) (model.User, error)
	FindBy(field string, value interface{}) (model.User, error)
}
