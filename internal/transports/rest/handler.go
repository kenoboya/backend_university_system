package rest

import (
	"errors"
	"fmt"
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
	Faculties    Faculties
	Specialties  Specialties
	Admins       Admins
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.Manager) *Handler {
	return &Handler{
		tokenManager: &tokenManager,
		Students:     NewStudentsHandler(services.Students),
		Users:        NewUsersHandler(services.Users),
		Teachers:     NewTeachersHandler(services.Teachers),
		Employees:    NewEmployeesHandler(services.Employees),
		Admins:       NewAdminsHandler(*services),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))
	h.initUsersRoutes(router)
	h.initAdminsRoutes(router)
	h.initEmployeeRoutes(router)
	h.initStudentsRoutes(router)
	h.initTeachersRoutes(router)
	fmt.Println(router)
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
	AdminPeople
}
type AdminPeople interface {
	createPerson(w http.ResponseWriter, r *http.Request)
	getPeople(w http.ResponseWriter, r *http.Request)
	getPerson(w http.ResponseWriter, r *http.Request)
	updatePerson(w http.ResponseWriter, r *http.Request)
	deletePerson(w http.ResponseWriter, r *http.Request)
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

type Specialties interface {
	// getSpecialties(w http.ResponseWriter, r *http.Request)
	// getSpecialty(w http.ResponseWriter, r *http.Request)
}
type Faculties interface {
	// getFaculties(w http.ResponseWriter, r *http.Request)
	// getFaculty(w http.ResponseWriter, r *http.Request)
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
