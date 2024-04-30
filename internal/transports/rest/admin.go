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
	admin := router.PathPrefix("/admin").Subrouter()
	{
		admin.Use()
		//admin.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		// admin.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)
		admin.HandleFunc("", h.createStudent).Methods(http.MethodPost)
		admin.HandleFunc("", h.getStudents).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getStudent).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.updateStudent).Methods(http.MethodPatch)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteStudent).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createTeacher).Methods(http.MethodPost)
		admin.HandleFunc("", h.getTeachers).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getTeacher).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.updateTeacher).Methods(http.MethodPatch)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteTeacher).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createEmployee).Methods(http.MethodPost)
		admin.HandleFunc("", h.getEmployees).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getEmployee).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.updateEmployee).Methods(http.MethodPatch)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteEmployee).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createSubject).Methods(http.MethodPost)
		admin.HandleFunc("", h.getSubject).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getSubject).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.updateSubject).Methods(http.MethodPatch)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteSubject).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createLesson).Methods(http.MethodPost)
		admin.HandleFunc("", h.getLessons).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getLesson).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteLesson).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createFaculty).Methods(http.MethodPost)
		admin.HandleFunc("", h.getFaculties).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getFaculty).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteFaculty).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createSpecialty).Methods(http.MethodPost)
		admin.HandleFunc("", h.getSpecialty).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getSpecialty).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.updateSpecialty).Methods(http.MethodPatch)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteSpecialty).Methods(http.MethodDelete)

		admin.HandleFunc("", h.createGroup).Methods(http.MethodPost)
		admin.HandleFunc("", h.getGroups).Methods(http.MethodGet)
		admin.HandleFunc("/{id:[0-9]+}", h.getGroups).Methods(http.MethodPost)
		admin.HandleFunc("/{id:[0-9]+}", h.deleteGroup).Methods(http.MethodDelete)
	}
}
