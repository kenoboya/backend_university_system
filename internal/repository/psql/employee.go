package psql

import "github.com/jmoiron/sqlx"

type EmployeesRepository struct {
	db *sqlx.DB
}

func NewEmployeesRepository(db *sqlx.DB) *EmployeesRepository {
	return &EmployeesRepository{db}
}
