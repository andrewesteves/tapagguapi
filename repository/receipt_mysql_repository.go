package repository

import (
	"database/sql"

	"github.com/andrewesteves/tapagguapi/model"
)

type ReceiptMySQLRepository struct {
	Conn *sql.DB
}

func NewReceiptMySQLRepository(Conn *sql.DB) ReceiptContractRepository {
	return &ReceiptMySQLRepository{Conn}
}

func (r ReceiptMySQLRepository) All() ([]model.Receipt, error) {
	return []model.Receipt{}, nil
}

func (r ReceiptMySQLRepository) Find() (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptMySQLRepository) Store(receipt model.Receipt) (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptMySQLRepository) Update(receipt model.Receipt) (model.Receipt, error) {
	return model.Receipt{}, nil
}

func (r ReceiptMySQLRepository) Destroy(id int64) (model.Receipt, error) {
	return model.Receipt{}, nil
}
