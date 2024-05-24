package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

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

func (r SubjectsRepository) Create(ctx context.Context, subject model.CreateSubjectInput) error {
	_, err := r.db.NamedExec("INSERT INTO subjects(name, semester, subject_type) VALUES(:name, :semester, :subject_type)", subject)
	if err != nil {
		return err
	}
	return nil
}
func (r SubjectsRepository) GetAll(ctx context.Context) ([]model.Subject, error) {
	subjects := []model.Subject{}
	err := r.db.Select(&subjects, "SELECT * FROM subjects")
	if err != nil {
		return subjects, err
	}
	return subjects, nil
}
func (r SubjectsRepository) GetById(ctx context.Context, id uint64) (model.Subject, error) {
	var subject model.Subject
	err := r.db.Get(&subject, "SELECT * FROM subjects WHERE subject_id = $1", id)
	if err != nil {
		return subject, err
	}
	return subject, nil
}
func (r SubjectsRepository) Update(ctx context.Context, id uint64, subject model.UpdateSubjectInput) error {
	_, err := r.db.Exec("UPDATE subjects SET name = $1, semester = $2, subject_type = $3 WHERE subject_id = $4",
		subject.Name, subject.Semester, subject.SubjectType, id)
	if err != nil {
		return err
	}
	return nil
}
func (r SubjectsRepository) Delete(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("DELETE FROM subjects WHERE subject_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r SubjectsRepository) GetSubjectsByStudentID(ctx context.Context, student_id uint64) ([]model.Subject, error) {
	subjects := []model.Subject{}
	err := r.db.Select(&subjects, "SELECT subject_id, name, semester, subject_type FROM subjects JOIN students_subjects USING(subject_id) WHERE student_id = $1", student_id)
	if err != nil {
		return subjects, err
	}
	return subjects, nil
}

func (r LessonsRepository) Create(ctx context.Context, lesson model.Lesson) error {
	_, err := r.db.NamedExec("INSERT INTO lessons(teacher_id, subject_id, lecture_room, time_start, time_end, lesson_type) VALUES(:teacher_id, :subject_id, :lecture_room, :time_start, :time_end, :lesson_type)", lesson)
	if err != nil {
		return err
	}
	return nil
}
func (r LessonsRepository) GetAll(ctx context.Context) ([]model.Lesson, error) {
	lessons := []model.Lesson{}
	err := r.db.Select(&lessons, "SELECT * FROM lessons")
	if err != nil {
		return lessons, err
	}
	return lessons, nil
}
func (r LessonsRepository) GetById(ctx context.Context, id uint64) (model.Lesson, error) {
	var lesson model.Lesson
	err := r.db.Get(&lesson, "SELECT * FROM lessons WHERE lesson_id = $1", id)
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}
func (r LessonsRepository) Delete(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("DELETE FROM lessons WHERE lesson_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r LessonsRepository) GetLessonsByStudentID(ctx context.Context, student_id uint64) ([]model.Lesson, error) {
	lessons := []model.Lesson{}
	err := r.db.Select(&lessons, "SELECT lesson_id, subject_id, teacher_id, lecture_room, time_start, time_end, lesson_type FROM lessons JOIN subjects USING (subject_id) JOIN students_subjects USING(subject_id) WHERE student_id = $1", student_id)
	if err != nil {
		return lessons, err
	}
	return lessons, nil
}
func (r LessonsRepository) GetLessonsByTeacherID(ctx context.Context, teacher_id uint64) ([]model.Lesson, error) {
	lessons := []model.Lesson{}
	err := r.db.Select(&lessons, "SELECT lesson_id, subject_id, teacher_id, lecture_room, time_start, time_end, lesson_type FROM lessons WHERE teacher_id = $1", teacher_id)
	if err != nil {
		return lessons, err
	}
	return lessons, nil
}
