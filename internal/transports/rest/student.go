package rest

import (
	"test-crud/internal/service"
)

type StudentsHandler struct {
	service service.Students
}

func NewStudentsHandler(service service.Students) *StudentsHandler {
	return &StudentsHandler{service: service}
}
