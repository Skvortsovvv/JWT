package store

import (
	"JWT/pkg/model"
	"github.com/google/uuid"
)

type Store interface {
	RegisterUser(user *model.User) error
	GetUser(login, password string) (uuid.UUID, error)
}
