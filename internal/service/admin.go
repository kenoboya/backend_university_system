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
	if response.Blocked {
		return s.repo.BlockUser(ctx, response.UserID)
	}
	return nil
}

func (s *AdminsService) UnblockUser(ctx context.Context, userID uint64) error {
	return s.repo.UnblockUser(ctx, userID)
}
