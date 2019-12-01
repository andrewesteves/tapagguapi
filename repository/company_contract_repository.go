package repository

import "github.com/andrewesteves/tapagguapi/model"

// CompanyContractRepository contract
type CompanyContractRepository interface {
	All(user model.User) ([]model.Company, error)
	Find(id int64) (model.Company, error)
	Store(c model.Company) (model.Company, error)
	Update(c model.Company) (model.Company, error)
	Destroy(int int64) (model.Company, error)
	FindBy(c model.Company, field string, value interface{}) (model.Company, error)
}
