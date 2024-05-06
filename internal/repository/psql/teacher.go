package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type TeachersRepository struct {
	db *sqlx.DB
}

func NewTeachersRepository(db *sqlx.DB) *TeachersRepository {
	return &TeachersRepository{db}
}

func (r *TeachersRepository) Create(ctx context.Context, teacher model.CreateTeacherInput) error {
	_, err := r.db.NamedExec("INSERT INTO teachers(employee_id VALUES(:employee_id)", teacher)
	if err != nil {
		return err
	}
	return nil
}
func (r *TeachersRepository) GetAll(ctx context.Context) ([]model.Teacher, error) {
	teachers := []model.Teacher{}
	err := r.db.Select(&teachers, "SELECT * FROM teachers JOIN employees USING (employee_id)")
	if err != nil {
		return teachers, err
	}
	return teachers, nil
}

func (r *TeachersRepository) GetById(ctx context.Context, id int64) (model.Teacher, error) {
	var teacher model.Teacher
	err := r.db.Get(&teacher, "SELECT * FROM teachers JOIN employees USING(employee_id) WHERE teacher_id=$1", id)
	if err != nil {
		return teacher, err
	}
	return teacher, nil
}
func (r *TeachersRepository) GetTeacherBriefInfoById(ctx context.Context, id int64) (model.TeacherBriefInfo, error) {
	var teacher model.TeacherBriefInfo
	err := r.db.Get(&teacher, "SELECT * FROM teachers JOIN employees USING(employee_id) JOIN people USING(person_id) WHERE teacher_id=$1", id)
	if err != nil {
		return teacher, err
	}
	return teacher, nil
}

func (r *TeachersRepository) Update(ctx context.Context, id int64, teacher model.UpdateTeacherInput) error {
	_, err := r.db.Exec("UPDATE teachers SET employee_id=$1, WHERE teacher_id=$2", teacher.EmployeeID, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *TeachersRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM teachers WHERE teacher_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
