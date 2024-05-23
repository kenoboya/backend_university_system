package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type StudentsRepository struct {
	db *sqlx.DB
}

func NewStudentsRepository(db *sqlx.DB) *StudentsRepository {
	return &StudentsRepository{db}
}

func (r StudentsRepository) Create(ctx context.Context, student model.CreateStudentInput) error {
	_, err := r.db.NamedExec("INSERT INTO students(person_id, group_id) VALUES(:person_id, :group_id)", student)
	if err != nil {
		return err
	}
	return nil
}
func (r StudentsRepository) GetAll(ctx context.Context) ([]model.Student, error) {
	student := []model.Student{}
	err := r.db.Select(&student, "SELECT * FROM students JOIN people USING(person_id) JOIN groups USING(group_id)")
	if err != nil {
		return student, err
	}
	return student, nil
}
func (r StudentsRepository) GetById(ctx context.Context, id uint64) (model.Student, error) {
	var student model.Student
	err := r.db.Get(&student, "SELECT * FROM students JOIN people USING(person_id) JOIN groups USING(group_id) WHERE student_id = $1", id)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentsRepository) GetStudentBriefInfoById(ctx context.Context, id uint64) (model.StudentBriefInfo, error) {
	var student model.StudentBriefInfo
	err := r.db.Get(&student, "SELECT * FROM students JOIN people USING(person_id) JOIN groups USING(group_id) WHERE student_id = $1", id)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentsRepository) GetStudentFullInfoById(ctx context.Context, id uint64) (model.StudentFullInfo, error) {
	var student model.StudentFullInfo
	err := r.db.Get(&student, "SELECT * FROM students JOIN people USING(person_id) JOIN groups USING(group_id) WHERE student_id = $1", id)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r StudentsRepository) Update(ctx context.Context, id uint64, student model.UpdateStudentInput) error {
	_, err := r.db.Exec("UPDATE students SET group_id= $1, WHERE student_id = $2", student.GroupID, id)
	if err != nil {
		return err
	}
	return nil
}
func (r StudentsRepository) Delete(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("DELETE FROM students WHERE student_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r StudentsRepository) GetStudentsAttendanceByLessonID(ctx context.Context, lesson_id uint64) ([]model.AttendanceRecord, error) {
	attendanceRecords := []model.AttendanceRecord{}
	err := r.db.Select(&attendanceRecords, "SELECT student_id, name, surname, lesson_id, status FROM attendance_grades JOIN students USING(student_id) WHERE lesson_id = $1", lesson_id)
	if err != nil {
		return attendanceRecords, err
	}
	return attendanceRecords, nil
}
func (r StudentsRepository) GetStudentsGradesByLessonID(ctx context.Context, lesson_id uint64) ([]model.Grade, error) {
	grades := []model.Grade{}
	err := r.db.Select(&grades, "SELECT student_id, name, surname, lesson_id, grade FROM attendance_grades JOIN students USING(student_id) WHERE lesson_id = $1", lesson_id)
	if err != nil {
		return grades, err
	}
	return grades, nil
}
