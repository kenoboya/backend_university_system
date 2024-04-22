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

func (h *StudentsHandler) initRoutes(router *mux.Router) {
	students := router.PathPrefix("/students").Subrouter()
	{
		// todo
	}
}
