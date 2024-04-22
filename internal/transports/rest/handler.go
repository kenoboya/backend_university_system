package rest

import (
	"net/http"
	_ "test-crud/docs"
	"test-crud/internal/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Deps struct {
	services *service.Services
}
type Handler struct {
	Students    Students
	Users       Users
	Teachers    Teachers
	Employees   Employees
	Subjects    Subjects
	Lessons     Lessons
	Faculties   Faculties
	Specialties Specialties
	Groups      Groups
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		Students:    NewStudentsHandler(services.Students),
		Users:       NewUsersHandler(services.Users),
		Teachers:    NewTeachersHandler(services.Teachers),
		Employees:   NewEmployeesHandler(services.Employees),
		Subjects:    NewSubjectsHandler(services.Subjects),
		Lessons:     NewLessonsHandler(services.Lessons),
		Faculties:   NewFacultiesHandler(services.Faculties),
		Specialties: NewSpecialtiesHandler(services.Specialties),
		Groups:      NewGroupsHandler(services.Groups),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))

	auth := router.PathPrefix("/login").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	}

	students := router.PathPrefix("/students").Subrouter()
	{
		students.HandleFunc("", h.createStudent).Methods(http.MethodPost)
		students.HandleFunc("/{id:[0-9]+}", h.deleteStudent).Methods(http.MethodDelete)
		students.HandleFunc("/{id:[0-9]+}", h.updateStudent).Methods(http.MethodPatch)
		students.HandleFunc("", h.getAllStudents).Methods(http.MethodGet)
		students.HandleFunc("/{id:[0-9]+}", h.getStudentById).Methods(http.MethodGet)
	}
	return router
}

type Routers interface {
	initRoutes(router *mux.Router)
}
type Users interface {
	Routers
	signUp(w http.ResponseWriter, r *http.Request)
	signIn(w http.ResponseWriter, r *http.Request)
}
type Students interface {
	Routers
	// todo
}
type Teachers interface {
	Routers
	// todo
}
type Employees interface {
	Routers
	// todo
}
type Subjects interface {
	Routers
	// todo
}
type Lessons interface {
	Routers
	// todo
}
type Specialties interface {
	Routers
	// todo
}
type Faculties interface {
	Routers
	// todo
}
type Groups interface {
	Routers
	// todo
}
