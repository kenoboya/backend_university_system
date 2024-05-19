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

func (s *PeopleService) GetListApplications(ctx context.Context) ([]model.PersonApplication, error) {
	return s.repo.GetAllApplications(ctx)
}
func (s *PeopleService) GetApplication(ctx context.Context, personID int64) (model.PersonApplication, error) {
	return s.repo.GetApplication(ctx, personID)
}
func (s *PeopleService) ResponseToApplication(ctx context.Context, response model.PersonApplication) error {
	if err := s.repo.UpdateApplicationStatus(ctx, response.Accepted, response.ApplicationID); err != nil {
		return err
	}
	if !response.Accepted {
		return s.Delete(ctx, response.PersonID)
	}
	return nil
}

func (s *PeopleService) CreateApplicationPerson(ctx context.Context, input model.CreatePersonInput) error {
	person, err := s.repo.GetPersonByUserID(ctx, input.UserID)
	if err != nil {
		return err
	}
	return s.repo.CreateApplicationPerson(ctx, person.PersonID)
}
