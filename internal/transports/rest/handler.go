package rest

import (
	"errors"
	"net/http"
	"strconv"
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
	Admins      Admins
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		//Students:    NewStudentsHandler(services.Students),
		Users:       NewUsersHandler(services.Users),
		Teachers:    NewTeachersHandler(services.Teachers),
		Employees:   NewEmployeesHandler(services.Employees),
		Subjects:    NewSubjectsHandler(services.Subjects),
		Lessons:     NewLessonsHandler(services.Lessons),
		Faculties:   NewFacultiesHandler(services.Faculties),
		Specialties: NewSpecialtiesHandler(services.Specialties),
		Groups:      NewGroupsHandler(services.Groups),
		Admins:      NewAdminsHandler(*services),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))
	h.Users.initRoutes(router)
	// auth := router.PathPrefix("/login").Subrouter()
	// {
	// 	auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	// }

	// students := router.PathPrefix("/students").Subrouter()
	// {
	// 	students.HandleFunc("", h.createStudent).Methods(http.MethodPost)
	// 	students.HandleFunc("/{id:[0-9]+}", h.deleteStudent).Methods(http.MethodDelete)
	// 	students.HandleFunc("/{id:[0-9]+}", h.updateStudent).Methods(http.MethodPatch)
	// 	students.HandleFunc("", h.getAllStudents).Methods(http.MethodGet)
	// 	students.HandleFunc("/{id:[0-9]+}", h.getStudentById).Methods(http.MethodGet)
	// }
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
	// Crud
}
type Crud interface {
	create(w http.ResponseWriter, r *http.Request)
	getAll(w http.ResponseWriter, r *http.Request)
	getById(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
}
type Teachers interface {
	Routers
	// Crud
}
type Employees interface {
	Routers
	//Crud
}
type Subjects interface {
	Routers
	// Crud
}
type Lessons interface {
	Routers
	// Crud
}
type Specialties interface {
	Routers
	// Crud
}
type Faculties interface {
	Routers
	// Crud
}
type Groups interface {
	Routers
	// Crud
}
type Admins interface {
	createStudent(w http.ResponseWriter, r *http.Request)
	getStudents(w http.ResponseWriter, r *http.Request)
	getStudent(w http.ResponseWriter, r *http.Request)
	updateStudent(w http.ResponseWriter, r *http.Request)
	deleteStudent(w http.ResponseWriter, r *http.Request)

	createTeacher(w http.ResponseWriter, r *http.Request)
	getTeachers(w http.ResponseWriter, r *http.Request)
	getTeacher(w http.ResponseWriter, r *http.Request)
	updateTeacher(w http.ResponseWriter, r *http.Request)
	deleteTeacher(w http.ResponseWriter, r *http.Request)

	createLesson(w http.ResponseWriter, r *http.Request)
	getLessons(w http.ResponseWriter, r *http.Request)
	getLesson(w http.ResponseWriter, r *http.Request)
	updateLesson(w http.ResponseWriter, r *http.Request)
	deleteLesson(w http.ResponseWriter, r *http.Request)

	createFaculty(w http.ResponseWriter, r *http.Request)
	getFaculties(w http.ResponseWriter, r *http.Request)
	getFaculty(w http.ResponseWriter, r *http.Request)
	updateFaculty(w http.ResponseWriter, r *http.Request)
	deleteFaculty(w http.ResponseWriter, r *http.Request)

	createSpecialty(w http.ResponseWriter, r *http.Request)
	getSpecialties(w http.ResponseWriter, r *http.Request)
	getSpecialty(w http.ResponseWriter, r *http.Request)
	updateSpecialty(w http.ResponseWriter, r *http.Request)
	deleteSpecialty(w http.ResponseWriter, r *http.Request)

	createGroup(w http.ResponseWriter, r *http.Request)
	getGroups(w http.ResponseWriter, r *http.Request)
	getGroup(w http.ResponseWriter, r *http.Request)
	updateGroup(w http.ResponseWriter, r *http.Request)
	deleteGroup(w http.ResponseWriter, r *http.Request)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("Id couldn't be zero")
	}
	return id, nil
}
