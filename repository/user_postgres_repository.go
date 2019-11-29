package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/andrewesteves/tapagguapi/model"
)

type UserPostgresRepository struct {
	Conn *sql.DB
}

func NewUserPostgresRepository(Conn *sql.DB) UserContractRepository {
	return &UserPostgresRepository{Conn}
}

func (r UserPostgresRepository) All() ([]model.User, error) {
	var users []model.User
	rs, err := r.Conn.Query("SELECT id, name, email, token, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var user model.User
		err = rs.Scan(&user.ID, &user.Name, &user.Email, &user.Token, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r UserPostgresRepository) Find(id int64) (model.User, error) {
	var user model.User
	err := r.Conn.QueryRow("SELECT id, name, email, token, created_at, updated_at FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Token, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r UserPostgresRepository) Store(user model.User) (model.User, error) {
	lastInsertId := 0
	err := r.Conn.QueryRow("INSERT INTO users (name, email, password, token, created_at, updated_at) VALUES ($1,$2,$3,$4,now(),now()) RETURNING id", user.Name, user.Email, user.Password, user.Token).Scan(&lastInsertId)
	if err != nil {
		return model.User{}, err
	}
	user.ID = int64(lastInsertId)
	return user, nil
}

func (r UserPostgresRepository) Update(user model.User) (model.User, error) {
	u, err := r.Find(user.ID)
	if err != nil {
		return model.User{}, err
	}
	if u.ID < 1 {
		return model.User{}, errors.New("Cant't find this user id.")
	}
	rs, err := r.Conn.Prepare("UPDATE users SET name = $1, email = $2, password = $3, token = $4, updated_at = now() WHERE id = $5")
	if err != nil {
		return model.User{}, err
	}
	rs.Exec(user.Name, user.Email, user.Password, user.Token, user.ID)
	return user, nil
}

func (r UserPostgresRepository) Destroy(id int64) (model.User, error) {
	u, err := r.Find(id)
	if err != nil {
		return model.User{}, err
	}
	if u.ID < 1 {
		return model.User{}, errors.New("Cant't find this user id.")
	}
	rs, err := r.Conn.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return model.User{}, err
	}
	rs.Exec(id)
	return u, nil
}

func (r UserPostgresRepository) FindBy(field string, value interface{}) (model.User, error) {
	var user model.User
	err := r.Conn.QueryRow(fmt.Sprintf("SELECT id, name, email, password, token, created_at, updated_at FROM users WHERE %s = $1", field), value).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
