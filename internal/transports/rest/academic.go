package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type FacultiesHandler struct {
	service service.Faculties
}

func NewFacultiesHandler(service service.Faculties) *FacultiesHandler {
	return &FacultiesHandler{service: service}
}

func (h *FacultiesHandler) initRoutes(router *mux.Router) {
	// faculties := router.PathPrefix("/faculties").Subrouter()
	// {
	// 	// todo
	// }
}

type SpecialtiesHandler struct {
	service service.Specialties
}

func NewSpecialtiesHandler(service service.Specialties) *SpecialtiesHandler {
	return &SpecialtiesHandler{service: service}
}

func (h *SpecialtiesHandler) initRoutes(router *mux.Router) {
	// specialties := router.PathPrefix("/specialties").Subrouter()
	// {
	// 	// todo
	// }
}

type GroupsHandler struct {
	service service.Groups
}

func NewGroupsHandler(service service.Groups) *GroupsHandler {
	return &GroupsHandler{service: service}
}

func (h *GroupsHandler) initRoutes(router *mux.Router) {
	// groups := router.PathPrefix("/groups").Subrouter()
	// {
	// 	// todo
	// }
}
