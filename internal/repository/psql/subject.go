package psql

import "github.com/jmoiron/sqlx"

type SubjectsRepository struct {
	db *sqlx.DB
}

func NewSubjectsRepository(db *sqlx.DB) *SubjectsRepository {
	return &SubjectsRepository{db}
}

type LessonsRepository struct {
	db *sqlx.DB
}

func NewLessonsRepository(db *sqlx.DB) *LessonsRepository {
	return &LessonsRepository{db}
}
