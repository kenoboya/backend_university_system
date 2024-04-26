package psql

import (
	"github.com/jmoiron/sqlx"
)

type AdminsRepository struct {
	db *sqlx.DB
}

func NewAdminsRepository(db *sqlx.DB) *AdminsRepository {
	return &AdminsRepository{db}
}
