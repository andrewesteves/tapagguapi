package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/andrewesteves/tapagguapi/model"
)

// CompanyPostgresRepository struct
type CompanyPostgresRepository struct {
	Conn *sql.DB
}

// NewCompanyPostgresRepository new repository
func NewCompanyPostgresRepository(Conn *sql.DB) CompanyContractRepository {
	return &CompanyPostgresRepository{Conn}
}

// All companies
func (r CompanyPostgresRepository) All(user model.User) ([]model.Company, error) {
	var companies []model.Company
	rs, err := r.Conn.Query("SELECT id, cnpj, name, title FROM companies WHERE user_id = $1", user.ID)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var company model.Company
		err = rs.Scan(&company.ID, &company.CNPJ, &company.Name, &company.Title)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

// Find company
func (r CompanyPostgresRepository) Find(id int64) (model.Company, error) {
	var company model.Company
	err := r.Conn.QueryRow("SELECT id, cnpj, name, title FROM companies WHERE id = $1", id).Scan(&company.ID, &company.CNPJ, &company.Name, &company.Title)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}

// Store company
func (r CompanyPostgresRepository) Store(company model.Company) (model.Company, error) {
	lastInsertID := 0
	err := r.Conn.QueryRow("INSERT INTO companies (user_id, cnpj, name, title, created_at, updated_at) VALUES ($1,$2,$3,$4,now(),now()) RETURNING id", company.User.ID, company.CNPJ, company.Name, company.Title).Scan(&lastInsertID)
	if err != nil {
		return model.Company{}, err
	}
	company.ID = int64(lastInsertID)
	return company, nil
}

// Update company
func (r CompanyPostgresRepository) Update(company model.Company) (model.Company, error) {
	u, err := r.Find(company.ID)
	if err != nil {
		return model.Company{}, err
	}
	if u.ID < 1 {
		return model.Company{}, errors.New("Cant't find this company id")
	}
	rs, err := r.Conn.Prepare("UPDATE companies SET cnpj = $1, name = $2, title = $3, updated_at = now() WHERE id = $4")
	if err != nil {
		return model.Company{}, err
	}
	rs.Exec(company.CNPJ, company.Name, company.Title, company.ID)
	return company, nil
}

// Destroy company
func (r CompanyPostgresRepository) Destroy(id int64) (model.Company, error) {
	u, err := r.Find(id)
	if err != nil {
		return model.Company{}, err
	}
	if u.ID < 1 {
		return model.Company{}, errors.New("Cant't find this company id")
	}
	rs, err := r.Conn.Prepare("DELETE FROM companies WHERE id = $1")
	if err != nil {
		return model.Company{}, err
	}
	rs.Exec(id)
	return u, nil
}

// FindBy company
func (r CompanyPostgresRepository) FindBy(company model.Company, field string, value interface{}) (model.Company, error) {
	err := r.Conn.QueryRow(fmt.Sprintf("SELECT id, cnpj, name, title FROM companies WHERE %s = $1 AND user_id = $2", field), value, company.User.ID).Scan(&company.ID, &company.CNPJ, &company.Name, &company.Title)
	if err != nil {
		return model.Company{}, err
	}
	return company, nil
}
