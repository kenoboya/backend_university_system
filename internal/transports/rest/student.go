package rest

import "github.com/gorilla/mux"

func (h *Handler) initStudentsRoutes(router *mux.Router) {
	students := router.PathPrefix("/student").Subrouter()
	{
		hub := students.PathPrefix("/hub").Subrouter()
		{
			h.Students.InitStudentProfileRoutes(hub)
		}
	}
}
