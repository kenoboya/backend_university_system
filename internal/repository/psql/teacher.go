package psql

import "github.com/jmoiron/sqlx"

type TeachersRepository struct {
	db *sqlx.DB
}

func NewTeachersRepository(db *sqlx.DB) *TeachersRepository {
	return &TeachersRepository{db}
}
