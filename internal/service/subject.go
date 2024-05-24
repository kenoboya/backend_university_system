package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"time"
)

type SubjectsService struct {
	repo psql.Subjects
}

func NewSubjectsService(repo psql.Subjects) *SubjectsService {
	return &SubjectsService{repo}
}

type LessonsService struct {
	repo psql.Lessons
}

func NewLessonsService(repo psql.Lessons) *LessonsService {
	return &LessonsService{repo}
}

func (s *SubjectsService) Create(ctx context.Context, subject model.CreateSubjectInput) error {
	return s.repo.Create(ctx, subject)
}
func (s *SubjectsService) GetAll(ctx context.Context) ([]model.Subject, error) {
	return s.repo.GetAll(ctx)
}
func (s *SubjectsService) GetById(ctx context.Context, id uint64) (model.Subject, error) {
	return s.repo.GetById(ctx, id)
}
func (s *SubjectsService) Update(ctx context.Context, id uint64, subject model.UpdateSubjectInput) error {
	return s.repo.Update(ctx, id, subject)
}
func (s *SubjectsService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *SubjectsService) GetStudentSubjects(ctx context.Context, student model.Student) ([]model.Subject, error) {
	return s.repo.GetSubjectsByStudentID(ctx, student.StudentID)
}

func (s *LessonsService) Create(ctx context.Context, input model.CreateLessonInput) error {
	lesson := model.Lesson{
		SubjectID:   input.SubjectID,
		TeacherID:   input.TeacherID,
		LectureRoom: input.LectureRoom,
		TimeStart:   input.Date,
		TimeEnd:     input.Date.Add(time.Minute * 45),
		LessonType:  input.LessonType,
	}
	return s.repo.Create(ctx, lesson)
}
func (s *LessonsService) GetAll(ctx context.Context) ([]model.Lesson, error) {
	return s.repo.GetAll(ctx)
}
func (s *LessonsService) GetById(ctx context.Context, id uint64) (model.Lesson, error) {
	return s.repo.GetById(ctx, id)
}
func (s *LessonsService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}
func (s *LessonsService) StudentSchedule(ctx context.Context, student model.Student) ([]model.Lesson, error) {
	return s.repo.GetLessonsByStudentID(ctx, student.StudentID)
}

func (s *LessonsService) TeacherSchedule(ctx context.Context, teacher model.Teacher) ([]model.Lesson, error) {
	return s.repo.GetLessonsByTeacherID(ctx, teacher.TeacherID)
}
