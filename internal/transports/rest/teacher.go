package rest

import "github.com/gorilla/mux"

func (h *Handler) initTeachersRoutes(router *mux.Router) {
	teachers := router.PathPrefix("/teacher").Subrouter()
	{
		hub := teachers.PathPrefix("/hub").Subrouter()
		{
			h.Teachers.InitTeacherProfileRoutes(hub)
		}
	}
}
