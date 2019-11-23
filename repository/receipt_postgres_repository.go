package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/lib/pq"
)

type ReceiptPostgresRepository struct {
	Conn *sql.DB
}

func NewReceiptPostgresRepository(Conn *sql.DB) ReceiptContractRepository {
	return &ReceiptPostgresRepository{Conn}
}

func (r ReceiptPostgresRepository) All() ([]model.Receipt, error) {
	var receipts []model.Receipt
	rs, err := r.Conn.Query("SELECT id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at FROM receipts")
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var receipt model.Receipt
		err = rs.Scan(&receipt.ID, &receipt.Title, &receipt.Tax, &receipt.Discount, &receipt.Extra, &receipt.Total, &receipt.URL, &receipt.AccessKey, &receipt.IssuedAt, &receipt.CreatedAt, &receipt.UpdatedAt)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}

func (r ReceiptPostgresRepository) Find(id int64) (model.Receipt, error) {
	var receipt model.Receipt
	err := r.Conn.QueryRow("SELECT id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at FROM receipts WHERE id = $1", id).Scan(&receipt.ID, &receipt.Title, &receipt.Tax, &receipt.Discount, &receipt.Extra, &receipt.Total, &receipt.URL, &receipt.AccessKey, &receipt.IssuedAt, &receipt.CreatedAt, &receipt.UpdatedAt)
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

func (r ReceiptPostgresRepository) Store(receipt model.Receipt) (model.Receipt, error) {
	lastInsertId := 0
	err := r.Conn.QueryRow("INSERT INTO receipts (title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,now(),now(),now()) RETURNING id", receipt.Title, receipt.Tax, receipt.Discount, receipt.Extra, receipt.Total, receipt.URL, receipt.AccessKey).Scan(&lastInsertId)
	if err != nil {
		return model.Receipt{}, err
	}
	receipt.ID = int64(lastInsertId)

	if len(receipt.Items) > 0 {
		txn, err := r.Conn.Begin()
		if err != nil {
			return model.Receipt{}, err
		}
		stmt, err := txn.Prepare(pq.CopyIn("items", "receipt_id", "title", "price", "quantity", "total", "tax", "measure", "created_at", "updated_at"))
		if err != nil {
			return model.Receipt{}, err
		}
		for _, item := range receipt.Items {
			_, err = stmt.Exec(receipt.ID, item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure, time.Now(), time.Now())
			if err != nil {
				return model.Receipt{}, err
			}
		}
		_, err = stmt.Exec()
		if err != nil {
			return model.Receipt{}, err
		}
		err = stmt.Close()
		if err != nil {
			return model.Receipt{}, err
		}
		err = txn.Commit()
		if err != nil {
			return model.Receipt{}, err
		}
	}

	return receipt, nil
}

func (r ReceiptPostgresRepository) Update(receipt model.Receipt) (model.Receipt, error) {
	rcpt, err := r.Find(receipt.ID)
	if err != nil {
		return model.Receipt{}, err
	}
	if rcpt.ID < 1 {
		return model.Receipt{}, errors.New("Cant't find this receipt id.")
	}
	rs, err := r.Conn.Prepare("UPDATE receipts SET title = $1, tax = $2, discount = $3, extra = $4, total = $5, url = $6, access_key = $7, issued_at = $8, updated_at = now() WHERE id = $9")
	if err != nil {
		return model.Receipt{}, err
	}
	rs.Exec(receipt.Title, receipt.Tax, receipt.Discount, receipt.Extra, receipt.Total, receipt.URL, receipt.AccessKey, receipt.IssuedAt, receipt.ID)
	return receipt, nil
}

func (r ReceiptPostgresRepository) Destroy(id int64) (model.Receipt, error) {
	rcpt, err := r.Find(id)
	if err != nil {
		return model.Receipt{}, err
	}
	if rcpt.ID < 1 {
		return model.Receipt{}, errors.New("Cant't find this receipt id.")
	}
	rs, err := r.Conn.Prepare("DELETE FROM receipts WHERE id = $1")
	if err != nil {
		return model.Receipt{}, err
	}
	rs.Exec(id)
	return rcpt, nil
}
