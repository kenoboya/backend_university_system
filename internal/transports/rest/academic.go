package rest

import (
	"test-crud/internal/service"
)

type FacultiesHandler struct {
	service service.Faculties
}

func NewFacultiesHandler(service service.Faculties) *FacultiesHandler {
	return &FacultiesHandler{service: service}
}

type SpecialtiesHandler struct {
	service service.Specialties
}

func NewSpecialtiesHandler(service service.Specialties) *SpecialtiesHandler {
	return &SpecialtiesHandler{service: service}
}

type GroupsHandler struct {
	service service.Groups
}

func NewGroupsHandler(service service.Groups) *GroupsHandler {
	return &GroupsHandler{service: service}
}
