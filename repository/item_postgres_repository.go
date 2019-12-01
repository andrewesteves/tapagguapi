package repository

import (
	"database/sql"
	"errors"

	"github.com/andrewesteves/tapagguapi/model"
)

// ItemPostgresRepository struct
type ItemPostgresRepository struct {
	Conn *sql.DB
}

// NewItemPostgresRepository new repository
func NewItemPostgresRepository(Conn *sql.DB) ItemContractRepository {
	return &ItemPostgresRepository{Conn}
}

// All items from a receipt
func (r ItemPostgresRepository) All(receiptID int64) ([]model.Item, error) {
	var items []model.Item
	rs, err := r.Conn.Query("SELECT id, title, price, quantity, total, tax, measure, created_at, updated_at FROM items WHERE receipt_id = $1 ORDER BY title", receiptID)
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

// Find a item from a receipt
func (r ItemPostgresRepository) Find(id int64) (model.Item, error) {
	var item model.Item
	err := r.Conn.QueryRow("SELECT id, title, price, quantity, total, tax, measure, created_at, updated_at FROM items WHERE id = $1", id).Scan(&item.ID, &item.Title, &item.Price, &item.Quantity, &item.Total, &item.Tax, &item.Measure, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// Store an item from a receipt
func (r ItemPostgresRepository) Store(receiptID int64, item model.Item) (model.Item, error) {
	lastInsertID := 0
	err := r.Conn.QueryRow("INSERT INTO items (receipt_id, title, price, quantity, total, tax, measure, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,now(),now()) RETURNING id", receiptID, item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure).Scan(&lastInsertID)
	if err != nil {
		return model.Item{}, err
	}
	item.ID = int64(lastInsertID)
	return item, nil
}

// Update an item from a receipt
func (r ItemPostgresRepository) Update(item model.Item) (model.Item, error) {
	it, err := r.Find(item.ID)
	if err != nil {
		return model.Item{}, err
	}
	if it.ID < 1 {
		return model.Item{}, errors.New("Cant't find this item id")
	}
	rs, err := r.Conn.Prepare("UPDATE items SET title = $1, price = $2, quantity = $3, total = $4, tax = $5, measure = $6, updated_at = now() WHERE id = $7")
	if err != nil {
		return model.Item{}, err
	}
	rs.Exec(item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure, item.ID)
	return item, nil
}

// Destroy an item from a receipt
func (r ItemPostgresRepository) Destroy(id int64) (model.Item, error) {
	it, err := r.Find(id)
	if err != nil {
		return model.Item{}, err
	}
	if it.ID < 1 {
		return model.Item{}, errors.New("Cant't find this item id")
	}
	rs, err := r.Conn.Prepare("DELETE FROM items WHERE id = $1")
	if err != nil {
		return model.Item{}, err
	}
	rs.Exec(id)
	return it, nil
}
