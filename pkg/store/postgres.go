package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Open")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "db.Ping()")
	}

	return db, nil
}
