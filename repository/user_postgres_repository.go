package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/andrewesteves/tapagguapi/model"
)

// UserPostgresRepository struct
type UserPostgresRepository struct {
	Conn *sql.DB
}

// NewUserPostgresRepository new repository
func NewUserPostgresRepository(Conn *sql.DB) UserContractRepository {
	return &UserPostgresRepository{Conn}
}

// All users
func (r UserPostgresRepository) All() ([]model.User, error) {
	var users []model.User
	rs, err := r.Conn.Query("SELECT id, name, email, token, active, remember, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var user model.User
		err = rs.Scan(&user.ID, &user.Name, &user.Email, &user.Token, &user.Active, &user.Remember, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Find user
func (r UserPostgresRepository) Find(id int64) (model.User, error) {
	var user model.User
	err := r.Conn.QueryRow("SELECT id, name, email, password, token, active, remember, created_at, updated_at FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token, &user.Active, &user.Remember, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Store user
func (r UserPostgresRepository) Store(user model.User) (model.User, error) {
	lastInsertID := 0
	err := r.Conn.QueryRow("INSERT INTO users (name, email, password, token, active, remember, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,now(),now()) RETURNING id", user.Name, user.Email, user.Password, user.Token, user.Active, user.Remember).Scan(&lastInsertID)
	if err != nil {
		return model.User{}, err
	}
	user.ID = int64(lastInsertID)
	return user, nil
}

// Update user
func (r UserPostgresRepository) Update(user model.User) (model.User, error) {
	u, err := r.Find(user.ID)
	if err != nil {
		return model.User{}, err
	}
	if u.ID < 1 {
		return model.User{}, errors.New("Cant't find this user id")
	}
	rs, err := r.Conn.Prepare("UPDATE users SET name = $1, email = $2, password = $3, token = $4, active = $5, remember = $6, updated_at = now() WHERE id = $7")
	if err != nil {
		return model.User{}, err
	}
	rs.Exec(user.Name, user.Email, user.Password, user.Token, user.Active, user.Remember, user.ID)
	return user, nil
}

// Destroy user
func (r UserPostgresRepository) Destroy(id int64) (model.User, error) {
	u, err := r.Find(id)
	if err != nil {
		return model.User{}, err
	}
	if u.ID < 1 {
		return model.User{}, errors.New("Cant't find this user id")
	}
	rs, err := r.Conn.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return model.User{}, err
	}
	rs.Exec(id)
	return u, nil
}

// FindBy user by field name
func (r UserPostgresRepository) FindBy(field string, value interface{}) (model.User, error) {
	var user model.User
	err := r.Conn.QueryRow(fmt.Sprintf("SELECT id, name, email, password, token, active, remember, created_at, updated_at FROM users WHERE %s = $1", field), value).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token, &user.Active, &user.Remember, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// FindByArgs where args
func (r UserPostgresRepository) FindByArgs(args map[string]interface{}) (model.User, error) {
	var user model.User
	var fields []interface{}
	var values []interface{}
	i := 1
	query := "SELECT id, name, email, password, token, active, remember, created_at, updated_at FROM users WHERE "

	for k, v := range args {
		j := strconv.Itoa(i)
		query += "%s = $" + j + " AND "
		fields = append(fields, k)
		values = append(values, v)
		i++
	}
	query = strings.TrimSuffix(query, "AND ")
	err := r.Conn.QueryRow(fmt.Sprintf(query, fields...), values...).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token, &user.Active, &user.Remember, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
