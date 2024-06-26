package rest

import (
	"net/http"
	_ "test-crud/docs"
	"test-crud/internal/service"
	"test-crud/internal/transports/rest/admin"
	"test-crud/internal/transports/rest/employee"
	"test-crud/internal/transports/rest/guest"
	"test-crud/internal/transports/rest/student"
	"test-crud/internal/transports/rest/teacher"
	"test-crud/internal/transports/rest/user"
	"test-crud/pkg/auth"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	Students     Students
	Users        Users
	Teachers     Teachers
	Employees    Employees
	Guests       Guests
	Admins       Admins
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.Manager) *Handler {

	return &Handler{
		tokenManager: &tokenManager,
		Students:     student.NewStudentsHandler(services.Students, services.Teachers, services.Subjects, services.Lessons),
		Users:        user.NewUsersHandler(services.Users, services.Complaints, services.People),
		Teachers:     teacher.NewTeachersHandler(services.Students, services.Teachers, services.Lessons),
		Employees:    employee.NewEmployeesHandler(services.Employees),
		Admins:       admin.NewAdminsHandler(*services),
		Guests:       guest.NewGuestsHandler(services.Faculties, services.Specialties, services.News),
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))
	h.initGuestsRoutes(router)
	h.initUsersRoutes(router)
	h.initAdminsRoutes(router)
	h.initEmployeeRoutes(router)
	h.initStudentsRoutes(router)
	h.initTeachersRoutes(router)
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
	AdminComplaints
	AdminRoutes
	AdminApplicationPeople
}
type AdminRoutes interface {
	InitAdminPeopleRoutes(hubs *mux.Router)
	InitAdminTeachersRoutes(hubs *mux.Router)
	InitAdminStudentsRoutes(hubs *mux.Router)
	InitAdminEmployeesRoutes(hubs *mux.Router)
	InitAdminSubjectsRoutes(hubs *mux.Router)
	InitAdminLessonsRoutes(hubs *mux.Router)
	InitAdminFacultiesRoutes(hubs *mux.Router)
	InitAdminSpecialtiesRoutes(hubs *mux.Router)
	InitAdminGroupsRoutes(hubs *mux.Router)
	InitAdminComplaintsRoutes(hubs *mux.Router)
	InitAdminApplicationsRoutes(hubs *mux.Router)
	InitAdminPeopleRequestsRoutes(application *mux.Router)
}
type AdminComplaints interface {
	GetComplaints(w http.ResponseWriter, r *http.Request)
	GetComplaint(w http.ResponseWriter, r *http.Request)
	ResponseToComplaint(w http.ResponseWriter, r *http.Request)
}
type AdminApplicationPeople interface {
	GetPeopleApplications(w http.ResponseWriter, r *http.Request)
	GetPersonApplication(w http.ResponseWriter, r *http.Request)
	ResponseToApplication(w http.ResponseWriter, r *http.Request)
}
type AdminPeople interface {
	CreatePerson(w http.ResponseWriter, r *http.Request)
	GetPeople(w http.ResponseWriter, r *http.Request)
	GetPerson(w http.ResponseWriter, r *http.Request)
	UpdatePerson(w http.ResponseWriter, r *http.Request)
	DeletePerson(w http.ResponseWriter, r *http.Request)
}
type AdminStudents interface {
	CreateStudent(w http.ResponseWriter, r *http.Request)
	GetStudents(w http.ResponseWriter, r *http.Request)
	GetStudent(w http.ResponseWriter, r *http.Request)
	UpdateStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudent(w http.ResponseWriter, r *http.Request)
}
type AdminTeachers interface {
	CreateTeacher(w http.ResponseWriter, r *http.Request)
	GetTeachers(w http.ResponseWriter, r *http.Request)
	GetTeacher(w http.ResponseWriter, r *http.Request)
	UpdateTeacher(w http.ResponseWriter, r *http.Request)
	DeleteTeacher(w http.ResponseWriter, r *http.Request)
}
type AdminEmployees interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	GetEmployees(w http.ResponseWriter, r *http.Request)
	GetEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
	DeleteEmployee(w http.ResponseWriter, r *http.Request)
}
type AdminSubjects interface {
	CreateSubject(w http.ResponseWriter, r *http.Request)
	GetSubjects(w http.ResponseWriter, r *http.Request)
	GetSubject(w http.ResponseWriter, r *http.Request)
	UpdateSubject(w http.ResponseWriter, r *http.Request)
	DeleteSubject(w http.ResponseWriter, r *http.Request)
}
type AdminLessons interface {
	CreateLesson(w http.ResponseWriter, r *http.Request)
	GetLessons(w http.ResponseWriter, r *http.Request)
	GetLesson(w http.ResponseWriter, r *http.Request)
	DeleteLesson(w http.ResponseWriter, r *http.Request)
}
type AdminFaculties interface {
	CreateFaculty(w http.ResponseWriter, r *http.Request)
	DeleteFaculty(w http.ResponseWriter, r *http.Request)
}
type AdminSpecialties interface {
	CreateSpecialty(w http.ResponseWriter, r *http.Request)
	UpdateSpecialty(w http.ResponseWriter, r *http.Request)
	DeleteSpecialty(w http.ResponseWriter, r *http.Request)
}
type AdminGroups interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
	DeleteGroup(w http.ResponseWriter, r *http.Request)
}
type Users interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	UserRoutes
	SubmitComplaint(w http.ResponseWriter, r *http.Request)
	SubmitPerson(w http.ResponseWriter, r *http.Request)
}
type UserRoutes interface {
	InitUserComplaintsRoutes(hubs *mux.Router)
	InitUserPeopleRoutes(hubs *mux.Router)
}
type Students interface {
	Profile
	StudentRoutes
	StudentSubjects
	Schedule
}
type Profile interface {
	GetStudentProfile(w http.ResponseWriter, r *http.Request)
	GetTeacherProfile(w http.ResponseWriter, r *http.Request)
}
type StudentSubjects interface {
	GetStudentSubjects(w http.ResponseWriter, r *http.Request)
	GetStudentSubject(w http.ResponseWriter, r *http.Request)
}
type Schedule interface {
	Schedule(w http.ResponseWriter, r *http.Request)
}
type StudentRoutes interface {
	InitStudentProfileRoutes(hubs *mux.Router)
	InitStudentSubjectsRoutes(hubs *mux.Router)
	InitStudentScheduleRoutes(hubs *mux.Router)
}
type Teachers interface {
	Profile
	TeacherRoutes
	Schedule
	TeacherLessons
}
type TeacherRoutes interface {
	InitTeacherProfileRoutes(hubs *mux.Router)
	InitTeacherScheduleRoutes(hub *mux.Router)
	InitTeacherLessonsRoutes(hub *mux.Router)
}
type TeacherLessons interface {
	GetLesson(w http.ResponseWriter, r *http.Request)
	AttendanceList(w http.ResponseWriter, r *http.Request)
	MarkAttendance(w http.ResponseWriter, r *http.Request)
	GradeList(w http.ResponseWriter, r *http.Request)
	Evaluate(w http.ResponseWriter, r *http.Request)
}
type Employees interface {
	// DELETE?
}

type Guests interface {
	GuestRoutes
	GuestFaculties
	GuestSpecialties
	GuestNews
}

type GuestFaculties interface {
	GetFaculties(w http.ResponseWriter, r *http.Request)
	GetFaculty(w http.ResponseWriter, r *http.Request)
}

type GuestNews interface {
	GetListNews(w http.ResponseWriter, r *http.Request)
	GetNews(w http.ResponseWriter, r *http.Request)
}

type GuestSpecialties interface {
	GetSpecialties(w http.ResponseWriter, r *http.Request)
	GetSpecialty(w http.ResponseWriter, r *http.Request)
}

type GuestRoutes interface {
	InitGuestFacultiesRoutes(hubs *mux.Router) *mux.Router
	InitGuestSpecialtiesRoutes(hubs *mux.Router)
	InitGuestNewsRoutes(hubs *mux.Router)
}
