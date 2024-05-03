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

func (h *Handler) initEmployeeRoutes(router *mux.Router) {
	// employees := router.PathPrefix("/employee_panel").Subrouter()
	// {

	// }
}
