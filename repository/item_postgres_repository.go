package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/andrewesteves/tapagguapi/model"
)

type ItemPostgresRepository struct {
	Conn *sql.DB
}

func NewItemPostgresRepository(Conn *sql.DB) ItemContractRepository {
	ItemPostgresRepository{}.Create()
	return &ItemPostgresRepository{Conn}
}

func (r ItemPostgresRepository) Create() {
	if _, err := r.Conn.Exec("CREATE TABLE IF NOT EXISTS items(id SERIAL PRIMARY KEY, receipt_id INTEGER NOT NULL, title VARCHAR(255) NOT NULL, price REAL NULL, quantity REAL NULL, total REAL NULL, tax REAL NULL, measure VARCHAR(255) NOT NULL, created_at TIMESTAMP NULL, updated_at TIMESTAMP NULL, CONSTRAINT fk_receipt_item FOREIGN KEY (receipt_id) REFERENCES receipts(id))"); err != nil {
		log.Fatalf("Error creating database table: %q", err)
	}
}

func (r ItemPostgresRepository) All(receiptId int64) ([]model.Item, error) {
	var items []model.Item
	rs, err := r.Conn.Query("SELECT id, title, price, quantity, total, tax, measure, created_at, updated_at FROM items WHERE receipt_id = $1", receiptId)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var item model.Item
		err = rs.Scan(&item.ID, &item.Title, &item.Price, &item.Quantity, &item.Total, &item.Tax, &item.Measure, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r ItemPostgresRepository) Find(id int64) (model.Item, error) {
	var item model.Item
	err := r.Conn.QueryRow("SELECT id, title, price, quantity, total, tax, measure, created_at, updated_at FROM items WHERE id = $1", id).Scan(&item.ID, &item.Title, &item.Price, &item.Quantity, &item.Total, &item.Tax, &item.Measure, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r ItemPostgresRepository) Store(receiptId int64, item model.Item) (model.Item, error) {
	lastInsertId := 0
	err := r.Conn.QueryRow("INSERT INTO items (receipt_id, title, price, quantity, total, tax, measure, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,now(),now()) RETURNING id", receiptId, item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure).Scan(&lastInsertId)
	if err != nil {
		return model.Item{}, err
	}
	item.ID = int64(lastInsertId)
	return item, nil
}

func (r ItemPostgresRepository) Update(item model.Item) (model.Item, error) {
	it, err := r.Find(item.ID)
	if err != nil {
		return model.Item{}, err
	}
	if it.ID < 1 {
		return model.Item{}, errors.New("Cant't find this item id.")
	}
	rs, err := r.Conn.Prepare("UPDATE items SET title = $1, price = $2, quantity = $3, total = $4, tax = $5, measure = $6, updated_at = now() WHERE id = $7")
	if err != nil {
		return model.Item{}, err
	}
	rs.Exec(item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure, item.ID)
	return item, nil
}

func (r ItemPostgresRepository) Destroy(id int64) (model.Item, error) {
	it, err := r.Find(id)
	if err != nil {
		return model.Item{}, err
	}
	if it.ID < 1 {
		return model.Item{}, errors.New("Cant't find this item id.")
	}
	rs, err := r.Conn.Prepare("DELETE FROM items WHERE id = $1")
	if err != nil {
		return model.Item{}, err
	}
	rs.Exec(id)
	return it, nil
}
