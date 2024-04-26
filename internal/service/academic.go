package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
)

type FacultiesService struct {
	repo psql.Faculties
}

func NewFacultiesService(repo psql.Faculties) *FacultiesService {
	return &FacultiesService{repo}
}

type SpecialtiesService struct {
	repo psql.Specialties
}

func NewSpecialtiesService(repo psql.Specialties) *SpecialtiesService {
	return &SpecialtiesService{repo}
}

type GroupsService struct {
	repo psql.Groups
}

func NewGroupsService(repo psql.Groups) *GroupsService {
	return &GroupsService{repo}
}

func (s *FacultiesService) Create(ctx context.Context, faculty model.CreateFacultyInput) error {
	return s.repo.Create(ctx, faculty)
}
func (s *FacultiesService) GetAll(ctx context.Context) ([]model.Faculty, error) {
	return s.repo.GetAll(ctx)
}
func (s *FacultiesService) GetById(ctx context.Context, id int64) (model.Faculty, error) {
	return s.repo.GetById(ctx, id)
}
func (s *FacultiesService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *SpecialtiesService) Create(ctx context.Context, specialty model.CreateSpecialtyInput) error {
	return s.repo.Create(ctx, specialty)
}
func (s *SpecialtiesService) GetAll(ctx context.Context) ([]model.Specialty, error) {
	return s.repo.GetAll(ctx)
}
func (s *SpecialtiesService) GetById(ctx context.Context, id int64) (model.Specialty, error) {
	return s.repo.GetById(ctx, id)
}
func (s *SpecialtiesService) Update(ctx context.Context, id int64, specialty model.UpdateSpecialtyInput) error {
	return s.repo.Update(ctx, id, specialty)
}
func (s *SpecialtiesService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *GroupsService) Create(ctx context.Context, group model.CreateGroupInput) error {
	return s.repo.Create(ctx, group)
}
func (s *GroupsService) GetAll(ctx context.Context) ([]model.Group, error) {
	return s.repo.GetAll(ctx)
}
func (s *GroupsService) GetById(ctx context.Context, id int64) (model.Group, error) {
	return s.repo.GetById(ctx, id)
}
func (s *GroupsService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
