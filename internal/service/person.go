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
func (s *PeopleService) GetById(ctx context.Context, id uint64) (model.Person, error) {
	return s.repo.GetById(ctx, id)
}
func (s *PeopleService) Update(ctx context.Context, id uint64, person model.UpdatePersonInput) error {
	return s.repo.Update(ctx, id, person)
}
func (s *PeopleService) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *PeopleService) GetListApplications(ctx context.Context) ([]model.PersonApplication, error) {
	return s.repo.GetAllApplications(ctx)
}
func (s *PeopleService) GetApplicationByID(ctx context.Context, applicationID uint64) (model.PersonApplication, error) {
	return s.repo.GetApplicationByID(ctx, applicationID)
}
func (s *PeopleService) GetApplicationByUserID(ctx context.Context, userID uint64) ([]model.PersonApplication, error) {
	return s.repo.GetApplicationsByUserID(ctx, userID)
}
func (s *PeopleService) ResponseToApplication(ctx context.Context, response model.PersonApplication) error {
	if err := s.repo.UpdateApplicationStatus(ctx, response.Status, response.ApplicationID); err != nil {
		return err
	}
	if response.Status == model.Denied {
		person, err := s.repo.GetPersonByUserID(ctx, response.UserID)
		if err != nil {
			return err
		}
		return s.Delete(ctx, person.PersonID)
	}
	return nil
}

func (s *PeopleService) CreateApplicationPerson(ctx context.Context, input model.CreatePersonInput) error {
	if err := s.repo.Create(ctx, input); err != nil {
		return err
	}
	return s.repo.CreateApplicationPerson(ctx, input)
}
