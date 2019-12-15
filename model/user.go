package model

import "time"

// User model
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password"`
	Token     string    `json:"token,omitempty"`
	Active    int64     `json:"active"`
	Remember  string    `json:"remember"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
