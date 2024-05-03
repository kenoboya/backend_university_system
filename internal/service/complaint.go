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

func (s *ComplaintsService) Create(ctx context.Context, reportingUserID int64, complaint model.CreateComplaintInput) error {
	return s.repo.Create(ctx, model.Complaint{
		ReportingUserID: reportingUserID,
		ReportedUserID:  complaint.ReportedUserID,
		Cause:           complaint.Cause,
		Time:            time.Now(),
	})
}

func (s *ComplaintsService) GetAll(ctx context.Context) ([]model.Complaint, error) {
	return s.repo.GetAll(ctx)
}

func (s *ComplaintsService) GetById(ctx context.Context, id int64) (model.Complaint, error) {
	return s.repo.GetById(ctx, id)
}

func (s *ComplaintsService) Response(ctx context.Context, complaintID int64, response string) error {
	return s.repo.Response(ctx, complaintID, response)
}
