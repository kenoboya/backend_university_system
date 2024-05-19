package admin

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type AdminsHandler struct {
	services service.Services
}

func NewAdminsHandler(services service.Services) *AdminsHandler {
	return &AdminsHandler{services: services}
}

func (h *AdminsHandler) InitAdminApplicationsRoutes(hubs *mux.Router) {
	applications := hubs.PathPrefix("/applications").Subrouter()
	{
		h.InitAdminPeopleRequestsRoutes(applications)
		h.InitAdminComplaintsRoutes(applications)
	}
}
