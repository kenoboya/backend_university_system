package rest

import "github.com/gorilla/mux"

func (h *Handler) initStudentsRoutes(router *mux.Router) {
	students := router.PathPrefix("/student").Subrouter()
	{
		students.Use(h.authStudentMiddleware)
		hub := students.PathPrefix("/hub").Subrouter()
		{
			h.Students.InitStudentProfileRoutes(hub)
			h.Students.InitStudentSubjectsRoutes(hub)
			h.Students.InitStudentScheduleRoutes(hub)
		}
	}
}
