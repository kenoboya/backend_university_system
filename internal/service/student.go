package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type StudentsService struct {
	repo psql.Students
}

func NewStudentsService(repo psql.Students) *StudentsService {
	return &StudentsService{repo}
}

func (s StudentsService) Create(ctx context.Context, student model.CreateStudentInput) error {
	return s.repo.Create(ctx, student)
}
func (s StudentsService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
func (s StudentsService) Update(ctx context.Context, id int64, student model.UpdateStudentInput) error {
	return s.repo.Update(ctx, id, student)
}
func (s StudentsService) GetById(ctx context.Context, id int64) (model.Student, error) {
	return s.repo.GetById(ctx, id)
}
func (s StudentsService) GetAll(ctx context.Context) ([]model.Student, error) {
	return s.repo.GetAll(ctx)
}
func (s *StudentsService) GetStudentProfile(ctx context.Context, id int64) (model.StudentBriefInfo, error) {
	return s.repo.GetStudentBriefInfoById(ctx, id)
}
