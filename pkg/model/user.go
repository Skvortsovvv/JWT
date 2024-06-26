package model

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `db:"id"`
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"password" db:"password"`
	Permission int       `db:"permission"`
}

func NewUser(id uuid.UUID, login, password string, permission int) *User {
	return &User{
		ID:         id,
		Login:      login,
		Password:   password,
		Permission: permission,
	}
}
