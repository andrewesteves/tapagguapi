package transformer

import "github.com/andrewesteves/tapagguapi/model"

// CompanyTransformer struct
type CompanyTransformer struct {
	ID    int64  `json:"id"`
	CNPJ  string `json:"cnpj"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

// TransformOne company specified JSON
func (rf CompanyTransformer) TransformOne(company model.Company) CompanyTransformer {
	var newCompany CompanyTransformer
	newCompany.ID = company.ID
	newCompany.CNPJ = company.CNPJ
	newCompany.Title = company.Title
	newCompany.Name = company.Name
	return newCompany
}
