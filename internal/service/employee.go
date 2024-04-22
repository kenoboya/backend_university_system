package service

import "test-crud/internal/repository/psql"

type EmployeesService struct {
	repo psql.Employees
}

func NewEmployeesService(repo psql.Employees) *EmployeesService {
	return &EmployeesService{repo}
}
