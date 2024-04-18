package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Event: NewEventPostgres(db),
	}
}
