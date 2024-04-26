package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
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
func (s *SubjectsService) GetById(ctx context.Context, id int64) (model.Subject, error) {
	return s.repo.GetById(ctx, id)
}
func (s *SubjectsService) Update(ctx context.Context, id int64, subject model.UpdateSubjectInput) error {
	return s.repo.Update(ctx, id, subject)
}
func (s *SubjectsService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *LessonsService) Create(ctx context.Context, lesson model.CreateLessonInput) error {
	return s.repo.Create(ctx, lesson)
}
func (s *LessonsService) GetAll(ctx context.Context) ([]model.Lesson, error) {
	return s.repo.GetAll(ctx)
}
func (s *LessonsService) GetById(ctx context.Context, id int64) (model.Lesson, error) {
	return s.repo.GetById(ctx, id)
}
func (s *LessonsService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
