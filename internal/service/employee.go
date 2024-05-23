package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type EmployeesService struct {
	repo psql.Employees
}

func NewEmployeesService(repo psql.Employees) *EmployeesService {
	return &EmployeesService{repo}
}

func (s *EmployeesService) Create(ctx context.Context, employee model.CreateEmployeeInput) error {
	return s.repo.Create(ctx, employee)
}
func (s *EmployeesService) GetAll(ctx context.Context) ([]model.Employee, error) {
	return s.repo.GetAll(ctx)
}
func (s *EmployeesService) GetById(ctx context.Context, id uint64) (model.Employee, error) {
	return s.repo.GetById(ctx, id)
}
func (s *EmployeesService) Update(ctx context.Context, id uint64, employee model.UpdateEmployeeInput) error {
	return s.repo.Update(ctx, id, employee)
}
func (s *EmployeesService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}
