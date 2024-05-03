package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type TeachersHandler struct {
	service service.Teachers
}

func NewTeachersHandler(service service.Teachers) *TeachersHandler {
	return &TeachersHandler{service: service}
}

func (h *Handler) initTeachersRoutes(router *mux.Router) {
	// teachers := router.PathPrefix("/teacher_panels").Subrouter()
	// {

	// }
}
