package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type Students struct {
	db *sqlx.DB
}

func NewStudents(db *sqlx.DB) *Students {
	return &Students{db}
}

func (st *Students) Create(ctx context.Context, student model.CreateStudentInput) error {
	_, err := st.db.NamedExec("INSERT INTO students(name,surname,age,email,phone,registered_at) VALUES(:name, :surname,:age, :email, :phone, :registered_at)", student)
	if err != nil {
		return err
	}
	return nil
}
func (st *Students) GetById(ctx context.Context, id int64) (model.Student, error) {
	var student model.Student
	err := st.db.Get(&student, "SELECT * FROM students WHERE student_id = $1", id)
	if err != nil {
		return student, err
	}
	return student, nil
}
func (st *Students) GetAll(ctx context.Context) ([]model.Student, error) {
	student := []model.Student{}
	err := st.db.Select(&student, "SELECT * FROM students")
	if err != nil {
		return student, err
	}
	return student, nil
}
func (st *Students) Update(ctx context.Context, id int64, student model.UpdateStudentInput) error {
	student.ID = id
	_, err := st.db.NamedExec("UPDATE students SET name= :name, surname= :surname, age= :age, email= :email, phone= :phone WHERE student_id = :student_id", student)
	if err != nil {
		return err
	}
	return nil
}
func (st *Students) Delete(ctx context.Context, id int64) error {
	_, err := st.db.Exec("DELETE FROM students WHERE student_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
