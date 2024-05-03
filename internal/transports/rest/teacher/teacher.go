package teacher

import (
	"test-crud/internal/service"
	"test-crud/internal/transports/rest"

	"github.com/gorilla/mux"
)

type TeachersHandler struct {
	service service.Teachers
}

func NewTeachersHandler(service service.Teachers) *TeachersHandler {
	return &TeachersHandler{service: service}
}

func (h *rest.Handler) initTeachersRoutes(router *mux.Router) {
	// teachers := router.PathPrefix("/teacher_panels").Subrouter()
	// {

	// }
}
