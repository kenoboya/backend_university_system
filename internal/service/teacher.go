package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type TeachersService struct {
	repo psql.Teachers
}

func NewTeachersService(repo psql.Teachers) *TeachersService {
	return &TeachersService{repo}
}

func (s *TeachersService) Create(ctx context.Context, teacher model.CreateTeacherInput) error {
	return s.repo.Create(ctx, teacher)
}
func (s *TeachersService) GetAll(ctx context.Context) ([]model.Teacher, error) {
	return s.repo.GetAll(ctx)
}
func (s *TeachersService) GetById(ctx context.Context, id int64) (model.Teacher, error) {
	return s.repo.GetById(ctx, id)
}
func (s *TeachersService) Update(ctx context.Context, id int64, teacher model.UpdateTeacherInput) error {
	return s.repo.Update(ctx, id, teacher)
}
func (s *TeachersService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
func (s *TeachersService) GetTeacherProfile(ctx context.Context, id int64) (model.TeacherBriefInfo, error) {
	return s.repo.GetTeacherBriefInfoById(ctx, id)
}
func (s *TeachersService) GetExtendedTeacherProfile(ctx context.Context, id int64) (model.TeacherFullInfo, error) {
	return s.repo.GetTeacherFullInfoById(ctx, id)
}
