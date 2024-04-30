package rest

import (
	"test-crud/internal/service"
)

type TeachersHandler struct {
	service service.Teachers
}

func NewTeachersHandler(service service.Teachers) *TeachersHandler {
	return &TeachersHandler{service: service}
}
