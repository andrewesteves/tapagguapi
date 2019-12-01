package service

import "github.com/andrewesteves/tapagguapi/model"

// CompanyContractService contract
type CompanyContractService interface {
	All(user model.User) ([]model.Company, error)
	Find(id int64) (model.Company, error)
	Store(r model.Company) (model.Company, error)
	Update(r model.Company) (model.Company, error)
	Destroy(int int64) (model.Company, error)
}
