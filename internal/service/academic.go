package service

import "test-crud/internal/repository/psql"

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
