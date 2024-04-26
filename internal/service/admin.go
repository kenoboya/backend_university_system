package service

import (
	"test-crud/internal/repository/psql"
)

type AdminsService struct {
	repo psql.Admins
}

func NewAdminsService(repo psql.Admins) *AdminsService {
	return &AdminsService{repo}
}
