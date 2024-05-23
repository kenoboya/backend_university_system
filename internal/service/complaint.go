package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"time"
)

type ComplaintsService struct {
	repo psql.Complaints
}

func NewComplaintsService(repo psql.Complaints) *ComplaintsService {
	return &ComplaintsService{repo}
}

func (s *ComplaintsService) Create(ctx context.Context, complaint model.CreateComplaintInput) error {
	return s.repo.Create(ctx, model.Complaint{
		ReportingUserID: complaint.ReportingUserID,
		ReportedUserID:  complaint.ReportedUserID,
		Cause:           complaint.Cause,
		Time:            time.Now(),
	})
}

func (s *ComplaintsService) GetAll(ctx context.Context) ([]model.Complaint, error) {
	return s.repo.GetAll(ctx)
}

func (s *ComplaintsService) GetById(ctx context.Context, id uint64) (model.Complaint, error) {
	return s.repo.GetById(ctx, id)
}

func (s *ComplaintsService) Response(ctx context.Context, complaintID uint64, response model.ResponseComplaintInput) error {
	return s.repo.Response(ctx, complaintID, response.Response)
}
