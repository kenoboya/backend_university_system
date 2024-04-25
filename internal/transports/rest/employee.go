package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type EmployeesHandler struct {
	service service.Employees
}

func NewEmployeesHandler(service service.Employees) *EmployeesHandler {
	return &EmployeesHandler{service: service}
}

func (h *EmployeesHandler) initRoutes(router *mux.Router) {
	// employees := router.PathPrefix("/employees").Subrouter()
	// {
	// 	// todo
	// }
}
