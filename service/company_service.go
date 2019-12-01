package service

import (
	"github.com/andrewesteves/tapagguapi/model"
	"github.com/andrewesteves/tapagguapi/repository"
)

// CompanyService struct
type CompanyService struct {
	companyRepository repository.CompanyContractRepository
}

// NewCompanyService new company service
func NewCompanyService(rs repository.CompanyContractRepository) CompanyContractService {
	return &CompanyService{rs}
}

// All companies service
func (r CompanyService) All(user model.User) ([]model.Company, error) {
	companies, err := r.companyRepository.All(user)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

// Find company service
func (r CompanyService) Find(id int64) (model.Company, error) {
	company, err := r.companyRepository.Find(id)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

// Store company service
func (r CompanyService) Store(company model.Company) (model.Company, error) {
	company, err := r.companyRepository.Store(company)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

// Update company service
func (r CompanyService) Update(company model.Company) (model.Company, error) {
	company, err := r.companyRepository.Update(company)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

// Destroy company service
func (r CompanyService) Destroy(id int64) (model.Company, error) {
	company, err := r.companyRepository.Destroy(id)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}
