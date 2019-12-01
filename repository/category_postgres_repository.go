package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/andrewesteves/tapagguapi/model"
)

// CategoryPostgresRepository struct
type CategoryPostgresRepository struct {
	Conn *sql.DB
}

// NewCategoryPostgresRepository new repository
func NewCategoryPostgresRepository(Conn *sql.DB) CategoryContractRepository {
	return &CategoryPostgresRepository{Conn}
}

// All categories
func (r CategoryPostgresRepository) All(user model.User) ([]model.Category, error) {
	var categories []model.Category
	rs, err := r.Conn.Query("SELECT id, title, icon FROM categories WHERE user_id = $1", user.ID)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var category model.Category
		err = rs.Scan(&category.ID, &category.Title, &category.Icon)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// Find category
func (r CategoryPostgresRepository) Find(id int64) (model.Category, error) {
	var category model.Category
	err := r.Conn.QueryRow("SELECT id, title, icon FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Title, &category.Icon)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

// Store category
func (r CategoryPostgresRepository) Store(category model.Category) (model.Category, error) {
	lastInsertID := 0
	err := r.Conn.QueryRow("INSERT INTO categories (user_id, title, icon, created_at, updated_at) VALUES ($1,$2,$3,now(),now()) RETURNING id", category.User.ID, category.Title, category.Icon).Scan(&lastInsertID)
	if err != nil {
		return model.Category{}, err
	}
	category.ID = int64(lastInsertID)
	return category, nil
}

// Update category
func (r CategoryPostgresRepository) Update(category model.Category) (model.Category, error) {
	u, err := r.Find(category.ID)
	if err != nil {
		return model.Category{}, err
	}
	if u.ID < 1 {
		return model.Category{}, errors.New("Cant't find this category id")
	}
	rs, err := r.Conn.Prepare("UPDATE categories SET title = $1, icon = $2, updated_at = now() WHERE id = $3")
	if err != nil {
		return model.Category{}, err
	}
	rs.Exec(category.Title, category.Icon, category.ID)
	return category, nil
}

// Destroy category
func (r CategoryPostgresRepository) Destroy(id int64) (model.Category, error) {
	u, err := r.Find(id)
	if err != nil {
		return model.Category{}, err
	}
	if u.ID < 1 {
		return model.Category{}, errors.New("Cant't find this category id")
	}
	rs, err := r.Conn.Prepare("DELETE FROM categories WHERE id = $1")
	if err != nil {
		return model.Category{}, err
	}
	rs.Exec(id)
	return u, nil
}

// FindBy category
func (r CategoryPostgresRepository) FindBy(category model.Category, field string, value interface{}) (model.Category, error) {
	err := r.Conn.QueryRow(fmt.Sprintf("SELECT id, title, icon FROM categories WHERE %s = $1 AND user_id = $2", field), value, category.User.ID).Scan(&category.ID, &category.Title, &category.Icon)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}
