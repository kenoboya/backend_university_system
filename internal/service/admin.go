package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type AdminsService struct {
	repo psql.Admins
}

func NewAdminsService(repo psql.Admins) *AdminsService {
	return &AdminsService{repo}
}

func (s *AdminsService) TryBlockUser(ctx context.Context, response model.ResponseComplaintInput) error {
	if response.Status {

	}
}

func (s *AdminsService) UnblockUser(ctx context.Context, userID int64) error {
	return s.repo.UnblockUser(ctx, userID)
}
