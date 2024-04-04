package store

import (
	"JWT/pkg/config"
	"JWT/pkg/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Store {
	return &Storage{
		db: db,
	}
}

func (s *Storage) RegisterUser(user *model.User) error {
	hashPass := config.GenerateHash(user.Password)
	user.Password = hashPass

	query := fmt.Sprint(`INSERT INTO users values($1, $2, $3, $4)`)

	_, err := s.db.Exec(query, user.ID, user.Login, user.Password, user.Permission)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUser(login, password string) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", "users")
	err := s.db.Get(&id, query, login, password)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *Storage) GetUserByID(id string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", "users")
	err := s.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
