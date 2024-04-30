package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type PeopleService struct {
	repo psql.People
}

func NewPeopleService(repo psql.People) *PeopleService {
	return &PeopleService{repo}
}

func (s *PeopleService) Create(ctx context.Context, person model.CreatePersonInput) error {
	return s.repo.Create(ctx, person)
}
func (s *PeopleService) GetAll(ctx context.Context) ([]model.Person, error) {
	return s.repo.GetAll(ctx)
}
func (s *PeopleService) GetById(ctx context.Context, id int64) (model.Person, error) {
	return s.repo.GetById(ctx, id)
}
func (s *PeopleService) Update(ctx context.Context, id int64, person model.UpdatePersonInput) error {
	return s.repo.Update(ctx, id, person)
}
func (s *PeopleService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
