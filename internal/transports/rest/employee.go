package rest

import (
	"test-crud/internal/service"
)

type EmployeesHandler struct {
	service service.Employees
}

func NewEmployeesHandler(service service.Employees) *EmployeesHandler {
	return &EmployeesHandler{service: service}
}
