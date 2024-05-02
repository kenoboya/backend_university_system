package rest

import (
	"errors"
	"net/http"
	"strconv"
	_ "test-crud/docs"
	"test-crud/internal/service"
	"test-crud/pkg/auth"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	Students     Students
	Users        Users
	Teachers     Teachers
	Employees    Employees
	Subjects     Subjects
	Lessons      Lessons
	Faculties    Faculties
	Specialties  Specialties
	Groups       Groups
	Admins       Admins
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.Manager) *Handler {
	return &Handler{
		tokenManager: &tokenManager,
		Students:     NewStudentsHandler(services.Students),
		Users:        NewUsersHandler(services.Users, services.People),
		Teachers:     NewTeachersHandler(services.Teachers),
		Employees:    NewEmployeesHandler(services.Employees),
		Subjects:     NewSubjectsHandler(services.Subjects),
		Lessons:      NewLessonsHandler(services.Lessons),
		Faculties:    NewFacultiesHandler(services.Faculties),
		Specialties:  NewSpecialtiesHandler(services.Specialties),
		Groups:       NewGroupsHandler(services.Groups),
		Admins:       NewAdminsHandler(*services),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))

	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("/sign-up", h.Users.signUp).Methods(http.MethodPost)
		users.HandleFunc("/sign-in", h.Users.signIn).Methods(http.MethodPost)
		users.HandleFunc("/refresh", h.Users.refresh).Methods(http.MethodGet)

		users.HandleFunc("", h.Users.createPerson).Methods(http.MethodPost)
		users.HandleFunc("", h.Users.getPeople).Methods(http.MethodGet)
		users.HandleFunc("/{id:[0-9]+}", h.Users.getPerson).Methods(http.MethodPost)
		users.HandleFunc("/{id:[0-9]+}", h.Users.updatePerson).Methods(http.MethodPatch)
		users.HandleFunc("/{id:[0-9]+}", h.Users.deletePerson).Methods(http.MethodDelete)
	}

	admin := router.PathPrefix("/admin").Subrouter()
	{
		admin.Use(h.authMiddleware)
		admin.HandleFunc("/students", h.Admins.createStudent).Methods(http.MethodPost)
		admin.HandleFunc("/students", h.Admins.getStudents).Methods(http.MethodGet)
		admin.HandleFunc("/students/{id:[0-9]+}", h.Admins.getStudent).Methods(http.MethodPost)
		admin.HandleFunc("/students/{id:[0-9]+}", h.Admins.updateStudent).Methods(http.MethodPatch)
		admin.HandleFunc("/students/{id:[0-9]+}", h.Admins.deleteStudent).Methods(http.MethodDelete)

		admin.HandleFunc("/teachers", h.Admins.createTeacher).Methods(http.MethodPost)
		admin.HandleFunc("/teachers", h.Admins.getTeachers).Methods(http.MethodGet)
		admin.HandleFunc("/teachers/{id:[0-9]+}", h.Admins.getTeacher).Methods(http.MethodGet)
		admin.HandleFunc("/teachers/{id:[0-9]+}", h.Admins.updateTeacher).Methods(http.MethodPatch)
		admin.HandleFunc("/teachers/{id:[0-9]+}", h.Admins.deleteTeacher).Methods(http.MethodDelete)

		admin.HandleFunc("/employees", h.Admins.createEmployee).Methods(http.MethodPost)
		admin.HandleFunc("/employees", h.Admins.getEmployees).Methods(http.MethodGet)
		admin.HandleFunc("/employees/{id:[0-9]+}", h.Admins.getEmployee).Methods(http.MethodGet)
		admin.HandleFunc("/employees/{id:[0-9]+}", h.Admins.updateEmployee).Methods(http.MethodPatch)
		admin.HandleFunc("/employees/{id:[0-9]+}", h.Admins.deleteEmployee).Methods(http.MethodDelete)

		admin.HandleFunc("/subjects", h.Admins.createSubject).Methods(http.MethodPost)
		admin.HandleFunc("/subjects", h.Admins.getSubject).Methods(http.MethodGet)
		admin.HandleFunc("/subjects/{id:[0-9]+}", h.Admins.getSubject).Methods(http.MethodGet)
		admin.HandleFunc("/subjects/{id:[0-9]+}", h.Admins.updateSubject).Methods(http.MethodPatch)
		admin.HandleFunc("/subjects/{id:[0-9]+}", h.Admins.deleteSubject).Methods(http.MethodDelete)

		admin.HandleFunc("/lessons", h.Admins.createLesson).Methods(http.MethodPost)
		admin.HandleFunc("/lessons", h.Admins.getLessons).Methods(http.MethodGet)
		admin.HandleFunc("/lessons/{id:[0-9]+}", h.Admins.getLesson).Methods(http.MethodGet)
		admin.HandleFunc("/lessons/{id:[0-9]+}", h.Admins.deleteLesson).Methods(http.MethodDelete)

		admin.HandleFunc("/faculties", h.Admins.createFaculty).Methods(http.MethodPost)
		admin.HandleFunc("/faculties", h.Admins.getFaculties).Methods(http.MethodGet)
		admin.HandleFunc("/faculties/{id:[0-9]+}", h.Admins.getFaculty).Methods(http.MethodGet)
		admin.HandleFunc("/faculties/{id:[0-9]+}", h.Admins.deleteFaculty).Methods(http.MethodDelete)

		admin.HandleFunc("/specialties", h.Admins.createSpecialty).Methods(http.MethodPost)
		admin.HandleFunc("/specialties", h.Admins.getSpecialty).Methods(http.MethodGet)
		admin.HandleFunc("/specialties/{id:[0-9]+}", h.Admins.getSpecialty).Methods(http.MethodGet)
		admin.HandleFunc("/specialties/{id:[0-9]+}", h.Admins.updateSpecialty).Methods(http.MethodPatch)
		admin.HandleFunc("/specialties/{id:[0-9]+}", h.Admins.deleteSpecialty).Methods(http.MethodDelete)

		admin.HandleFunc("/groups", h.Admins.createGroup).Methods(http.MethodPost)
		admin.HandleFunc("/groups", h.Admins.getGroups).Methods(http.MethodGet)
		admin.HandleFunc("/groups/{id:[0-9]+}", h.Admins.getGroups).Methods(http.MethodGet)
		admin.HandleFunc("/groups/{id:[0-9]+}", h.Admins.deleteGroup).Methods(http.MethodDelete)
	}
	return router
}

type Admins interface {
	AdminStudents
	AdminTeachers
	AdminEmployees
	AdminSubjects
	AdminLessons
	AdminFaculties
	AdminSpecialties
	AdminGroups
}
type AdminStudents interface {
	createStudent(w http.ResponseWriter, r *http.Request)
	getStudents(w http.ResponseWriter, r *http.Request)
	getStudent(w http.ResponseWriter, r *http.Request)
	updateStudent(w http.ResponseWriter, r *http.Request)
	deleteStudent(w http.ResponseWriter, r *http.Request)
}
type AdminTeachers interface {
	createTeacher(w http.ResponseWriter, r *http.Request)
	getTeachers(w http.ResponseWriter, r *http.Request)
	getTeacher(w http.ResponseWriter, r *http.Request)
	updateTeacher(w http.ResponseWriter, r *http.Request)
	deleteTeacher(w http.ResponseWriter, r *http.Request)
}
type AdminEmployees interface {
	createEmployee(w http.ResponseWriter, r *http.Request)
	getEmployees(w http.ResponseWriter, r *http.Request)
	getEmployee(w http.ResponseWriter, r *http.Request)
	updateEmployee(w http.ResponseWriter, r *http.Request)
	deleteEmployee(w http.ResponseWriter, r *http.Request)
}
type AdminSubjects interface {
	createSubject(w http.ResponseWriter, r *http.Request)
	getSubjects(w http.ResponseWriter, r *http.Request)
	getSubject(w http.ResponseWriter, r *http.Request)
	updateSubject(w http.ResponseWriter, r *http.Request)
	deleteSubject(w http.ResponseWriter, r *http.Request)
}
type AdminLessons interface {
	createLesson(w http.ResponseWriter, r *http.Request)
	getLessons(w http.ResponseWriter, r *http.Request)
	getLesson(w http.ResponseWriter, r *http.Request)
	deleteLesson(w http.ResponseWriter, r *http.Request)
}
type AdminFaculties interface {
	createFaculty(w http.ResponseWriter, r *http.Request)
	getFaculties(w http.ResponseWriter, r *http.Request)
	getFaculty(w http.ResponseWriter, r *http.Request)
	deleteFaculty(w http.ResponseWriter, r *http.Request)
}
type AdminSpecialties interface {
	createSpecialty(w http.ResponseWriter, r *http.Request)
	getSpecialties(w http.ResponseWriter, r *http.Request)
	getSpecialty(w http.ResponseWriter, r *http.Request)
	updateSpecialty(w http.ResponseWriter, r *http.Request)
	deleteSpecialty(w http.ResponseWriter, r *http.Request)
}
type AdminGroups interface {
	createGroup(w http.ResponseWriter, r *http.Request)
	getGroups(w http.ResponseWriter, r *http.Request)
	getGroup(w http.ResponseWriter, r *http.Request)
	deleteGroup(w http.ResponseWriter, r *http.Request)
}
type Users interface {
	signUp(w http.ResponseWriter, r *http.Request)
	signIn(w http.ResponseWriter, r *http.Request)
	refresh(w http.ResponseWriter, r *http.Request)
	UserPeople
}
type UserPeople interface {
	createPerson(w http.ResponseWriter, r *http.Request)
	getPeople(w http.ResponseWriter, r *http.Request)
	getPerson(w http.ResponseWriter, r *http.Request)
	updatePerson(w http.ResponseWriter, r *http.Request)
	deletePerson(w http.ResponseWriter, r *http.Request)
}

type Students interface {
	// DELETE?
}
type Teachers interface {
	// getTeachers(w http.ResponseWriter, r *http.Request)
	// getTeacher(w http.ResponseWriter, r *http.Request)
}
type Employees interface {
	// DELETE?
}
type Subjects interface {
	// getSubjects(w http.ResponseWriter, r *http.Request)
	// getSubject(w http.ResponseWriter, r *http.Request)
}
type Lessons interface {
	// DELETE?
}
type Specialties interface {
	// getSpecialties(w http.ResponseWriter, r *http.Request)
	// getSpecialty(w http.ResponseWriter, r *http.Request)
}
type Faculties interface {
	// getFaculties(w http.ResponseWriter, r *http.Request)
	// getFaculty(w http.ResponseWriter, r *http.Request)
}
type Groups interface {
	// DELETE?
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("id couldn't be zero")
	}
	return id, nil
}
