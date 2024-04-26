package rest

import (
	"net/http"
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type AdminsHandler struct {
	services service.Services
}

func NewAdminsHandler(services service.Services) *AdminsHandler {
	return &AdminsHandler{services: services}
}

func (h *AdminsHandler) initRoutes(router *mux.Router) {
	admins := router.PathPrefix("/admins").Subrouter()
	{
		admins.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		admins.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

		teachers := router.PathPrefix("/teachers").Subrouter()
		{
			teachers.HandleFunc("", h.createTeacher).Methods(http.MethodPost)
			teachers.HandleFunc("", h.getTeachers).Methods(http.MethodGet)
			teachers.HandleFunc("/{id:[0-9]+}", h.getTeacher).Methods(http.MethodPost)
			teachers.HandleFunc("/{id:[0-9]+}", h.updateTeacher).Methods(http.MethodPatch)
			teachers.HandleFunc("/{id:[0-9]+}", h.deleteTeacher).Methods(http.MethodDelete)
		}

		students := router.PathPrefix("/students").Subrouter()
		{
			students.HandleFunc("", h.createStudent).Methods(http.MethodPost)
			students.HandleFunc("", h.getStudents).Methods(http.MethodGet)
			students.HandleFunc("/{id:[0-9]+}", h.getStudent).Methods(http.MethodPost)
			students.HandleFunc("/{id:[0-9]+}", h.updateStudent).Methods(http.MethodPatch)
			students.HandleFunc("/{id:[0-9]+}", h.deleteStudent).Methods(http.MethodDelete)
		}
	}
}
