package transformer

import (
	"strings"

	"github.com/andrewesteves/tapagguapi/model"
)

// CompanyTransformer struct
type CompanyTransformer struct {
	ID      int64  `json:"id"`
	CNPJ    string `json:"cnpj"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Address string `json:"address"`
}

// TransformOne company specified JSON
func (rf CompanyTransformer) TransformOne(company model.Company) CompanyTransformer {
	var newCompany CompanyTransformer
	newCompany.ID = company.ID
	newCompany.CNPJ = company.CNPJ
	newCompany.Title = company.Title
	newCompany.Name = company.Name
	newCompany.Address = ""
	if company.Street != "" {
		newCompany.Address += company.Street + ", "
	}
	if company.Number != "" {
		newCompany.Address += company.Number + ", "
	}
	if company.District != "" {
		newCompany.Address += company.District + ", "
	}
	if company.City != "" {
		newCompany.Address += company.City + ", "
	}
	if company.State != "" {
		newCompany.Address += company.State + ", "
	}
	if company.Zipcode != "" {
		newCompany.Address += company.Zipcode + ", "
	}
	newCompany.Address = strings.TrimSuffix(newCompany.Address, ", ")
	return newCompany
}
