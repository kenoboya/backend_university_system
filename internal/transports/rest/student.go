package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type StudentsHandler struct {
	service service.Students
}

func NewStudentsHandler(service service.Students) *StudentsHandler {
	return &StudentsHandler{service: service}
}

func (h *Handler) initStudentsRoutes(router *mux.Router) {
	// students := router.PathPrefix("/student_panel").Subrouter()
	// {

	// }
}
