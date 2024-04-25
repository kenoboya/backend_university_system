package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type TeachersHandler struct {
	service service.Teachers
}

func NewTeachersHandler(service service.Employees) *TeachersHandler {
	return &TeachersHandler{service: service}
}

func (h *TeachersHandler) initRoutes(router *mux.Router) {
	// teachers := router.PathPrefix("/teachers").Subrouter()
	// {
	// 	// todo
	// }
}
