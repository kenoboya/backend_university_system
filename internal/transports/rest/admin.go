package rest

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

func (h *Handler) initAdminsRoutes(router *mux.Router) {
	admin := router.PathPrefix("/admin").Subrouter()
	{
		admin.Use(h.authMiddleware)
		hubs := admin.PathPrefix("/hub").Subrouter()
		{
			h.initAdminPeopleRoutes(hubs)
			h.initAdminTeachersRoutes(hubs)
			h.initAdminStudentsRoutes(hubs)
			h.initAdminEmployeesRoutes(hubs)
			h.initAdminSubjectsRoutes(hubs)
			h.initAdminLessonsRoutes(hubs)
			h.initAdminFacultiesRoutes(hubs)
			h.initAdminSpecialtiesRoutes(hubs)
			h.initAdminGroupsRoutes(hubs)

		}
		settings := admin.PathPrefix("/settings").Subrouter()
		{
			// Пока оставлю так, но в планах сделать настройки
			h.initAdminPeopleRoutes(settings)
		}
	}
}
