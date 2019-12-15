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

// GroupTotal categories grouped to get the total
func (r CategoryPostgresRepository) GroupTotal(user model.User, values map[string]string) ([]model.Category, error) {
	var categories []model.Category
	var rs *sql.Rows
	var err error
	args := []interface{}{user.ID}
	i := 2
	query := "SELECT c.id, c.title, c.icon, SUM(r.total) AS ctotal FROM categories AS c "
	query += "INNER JOIN receipts AS r "
	query += "ON c.id = r.category_id "
	query += "WHERE r.user_id = $1 "

	if len(values) == 0 {
		query += "AND (EXTRACT (MONTH FROM r.issued_at) = EXTRACT (MONTH FROM CURRENT_DATE)) AND (EXTRACT (YEAR FROM r.issued_at) = EXTRACT (YEAR FROM CURRENT_DATE)) "
	} else {
		if mt, ok := values["month"]; ok {
			args = append(args, mt)
			query += fmt.Sprintf("AND (EXTRACT (MONTH FROM r.issued_at) = $%d) ", i)
			i++
		}
		if yr, ok := values["year"]; ok {
			args = append(args, yr)
			query += fmt.Sprintf("AND (EXTRACT (YEAR FROM r.issued_at) = $%d) ", i)
		}
	}

	query += "GROUP BY c.id, c.title, c.icon "
	query += "ORDER BY ctotal DESC"

	rs, err = r.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var c model.Category
		err = rs.Scan(&c.ID, &c.Title, &c.Icon, &c.Total)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
