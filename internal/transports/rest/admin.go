package rest

import "github.com/gorilla/mux"

func (h *Handler) initAdminsRoutes(router *mux.Router) {
	admin := router.PathPrefix("/admin").Subrouter()
	{
		admin.Use(h.authAdminMiddleware)
		hubs := admin.PathPrefix("/hub").Subrouter()
		{
			h.Admins.InitAdminPeopleRoutes(hubs)
			h.Admins.InitAdminTeachersRoutes(hubs)
			h.Admins.InitAdminStudentsRoutes(hubs)
			h.Admins.InitAdminEmployeesRoutes(hubs)
			h.Admins.InitAdminSubjectsRoutes(hubs)
			h.Admins.InitAdminLessonsRoutes(hubs)
			h.Admins.InitAdminFacultiesRoutes(hubs)
			h.Admins.InitAdminSpecialtiesRoutes(hubs)
			h.Admins.InitAdminGroupsRoutes(hubs)
			// h.Admins.InitAdminComplaintsRoutes(hubs)
			h.Admins.InitAdminApplicationsRoutes(hubs)
		}
		settings := admin.PathPrefix("/settings").Subrouter()
		{
			// Пока оставлю так, но в планах сделать настройки
			h.Admins.InitAdminPeopleRoutes(settings)
		}
	}
}
