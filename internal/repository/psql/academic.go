package psql

import "github.com/jmoiron/sqlx"

type FacultiesRepository struct {
	db *sqlx.DB
}

func NewFacultiesRepository(db *sqlx.DB) *FacultiesRepository {
	return &FacultiesRepository{db}
}

type SpecialtiesRepository struct {
	db *sqlx.DB
}

func NewSpecialtiesRepository(db *sqlx.DB) *SpecialtiesRepository {
	return &SpecialtiesRepository{db}
}

type GroupsRepository struct {
	db *sqlx.DB
}

func NewGroupsRepository(db *sqlx.DB) *GroupsRepository {
	return &GroupsRepository{db}
}
