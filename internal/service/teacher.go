package service

import "test-crud/internal/repository/psql"

type TeachersService struct {
	repo psql.Teachers
}

func NewTeachersService(repo psql.Teachers) *TeachersService {
	return &TeachersService{repo}
}
