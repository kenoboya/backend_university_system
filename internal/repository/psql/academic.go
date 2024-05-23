package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

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

func (r FacultiesRepository) Create(ctx context.Context, faculty model.CreateFacultyInput) error {
	_, err := r.db.NamedExec("INSERT INTO faculties(faculty_id,full_name) VALUES(:faculty_id, :full_name)", faculty)
	if err != nil {
		return err
	}
	return nil
}
func (r FacultiesRepository) GetAll(ctx context.Context) ([]model.Faculty, error) {
	faculties := []model.Faculty{}
	err := r.db.Select(&faculties, "SELECT * FROM faculties")
	if err != nil {
		return faculties, err
	}
	return faculties, nil
}
func (r FacultiesRepository) GetById(ctx context.Context, id string) (model.Faculty, error) {
	var faculty model.Faculty
	err := r.db.Get(&faculty, "SELECT * FROM faculties WHERE faculty_id = $1", id)
	if err != nil {
		return faculty, err
	}
	return faculty, nil
}

func (r FacultiesRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM faculties WHERE faculty_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r SpecialtiesRepository) Create(ctx context.Context, specialty model.CreateSpecialtyInput) error {
	_, err := r.db.NamedExec("INSERT INTO specialties(specialty_id, faculty_id, full_name) VALUES(:specialty_id, :faculty_id, :full_name)", specialty)
	if err != nil {
		return err
	}
	return nil
}
func (r SpecialtiesRepository) GetAll(ctx context.Context) ([]model.Specialty, error) {
	specialties := []model.Specialty{}
	err := r.db.Select(&specialties, "SELECT * FROM specialties")
	if err != nil {
		return specialties, err
	}
	return specialties, nil
}
func (r SpecialtiesRepository) GetSpecialtiesByFacultyID(ctx context.Context, faculty_id string) ([]model.Specialty, error) {
	specialties := []model.Specialty{}
	err := r.db.Select(&specialties, "SELECT * FROM specialties WHERE faculty_id=$1", faculty_id)
	if err != nil {
		return specialties, err
	}
	return specialties, nil
}
func (r SpecialtiesRepository) GetById(ctx context.Context, id uint16) (model.Specialty, error) {
	var specialty model.Specialty
	err := r.db.Get(&specialty, "SELECT * FROM specialties WHERE specialty_id = $1", id)
	if err != nil {
		return specialty, err
	}
	return specialty, nil
}
func (r SpecialtiesRepository) Update(ctx context.Context, id uint16, specialty model.UpdateSpecialtyInput) error {
	_, err := r.db.Exec("UPDATE specialties SET faculty_id = $1, full_name = $2 WHERE employee_id = $3",
		specialty.FacultyID, specialty.FullName, id)
	if err != nil {
		return err
	}
	return nil
}
func (r SpecialtiesRepository) Delete(ctx context.Context, id uint16) error {
	_, err := r.db.Exec("DELETE FROM specialties WHERE specialty_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r GroupsRepository) Create(ctx context.Context, group model.Group) error {
	_, err := r.db.NamedExec("INSERT INTO groups(group_id,specialty_id, full_name, educational_level, start_year, end_year) VALUES(:group_id, :specialty_id, :full_name, :educational_level, :start_year, :end_year)", group)
	if err != nil {
		return err
	}
	return nil
}
func (r GroupsRepository) GetAll(ctx context.Context) ([]model.Group, error) {
	groups := []model.Group{}
	err := r.db.Select(&groups, "SELECT * FROM groups")
	if err != nil {
		return groups, err
	}
	return groups, nil
}
func (r GroupsRepository) GetById(ctx context.Context, id string) (model.Group, error) {
	var group model.Group
	err := r.db.Get(&group, "SELECT * FROM groups WHERE group_id = $1", id)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (r GroupsRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM groups WHERE group_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
