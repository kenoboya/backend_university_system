package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"test-crud/pkg/auth"
	"test-crud/pkg/hash"
	"time"
)

type Services struct {
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
	People      People
	Complaints  Complaints
	News        News
}
type Deps struct {
	Repos           psql.Repositories
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	return &Services{
		Students:    NewStudentsService(deps.Repos.Students),
		Users:       NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL),
		Teachers:    NewTeachersService(deps.Repos.Teachers),
		Employees:   NewEmployeesService(deps.Repos.Employees),
		Subjects:    NewSubjectsService(deps.Repos.Subjects),
		Lessons:     NewLessonsService(deps.Repos.Lessons),
		Faculties:   NewFacultiesService(deps.Repos.Faculties),
		Specialties: NewSpecialtiesService(deps.Repos.Specialties),
		Groups:      NewGroupsService(deps.Repos.Groups),
		Admins:      NewAdminsService(deps.Repos.Admins),
		People:      NewPeopleService(deps.Repos.People),
		Complaints:  NewComplaintsService(deps.Repos.Complaints),
		News:        NewNewsService(deps.Repos.News),
	}
}

type Students interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	GetAll(ctx context.Context) ([]model.Student, error)
	GetById(ctx context.Context, id int64) (model.Student, error)
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	Delete(ctx context.Context, id int64) error
	GetStudentProfile(ctx context.Context, id int64) (model.StudentBriefInfo, error)
	GetExtendedStudentProfile(ctx context.Context, id int64) (model.StudentFullInfo, error)
	GetStudentsAttendance(ctx context.Context, lesson_id int64) ([]model.AttendanceRecord, error)
	GetStudentsGrades(ctx context.Context, lesson_id int64) ([]model.Grade, error)
}
type Users interface {
	SignUp(ctx context.Context, input model.UserSignUpInput) error
	SignIn(ctx context.Context, input model.UserSignInInput) (Tokens, error)
	Refresh(ctx context.Context, refreshToken string) (Tokens, error)
	ChangeRole(ctx context.Context, role string, user_id int64) error
}
type Teachers interface {
	Create(ctx context.Context, teacher model.CreateTeacherInput) error
	GetAll(ctx context.Context) ([]model.Teacher, error)
	GetById(ctx context.Context, id int64) (model.Teacher, error)
	Update(ctx context.Context, id int64, teacher model.UpdateTeacherInput) error
	Delete(ctx context.Context, id int64) error
	GetTeacherProfile(ctx context.Context, id int64) (model.TeacherBriefInfo, error)
	GetExtendedTeacherProfile(ctx context.Context, id int64) (model.TeacherFullInfo, error)
	MarkAttendance(ctx context.Context, lesson_id int64, attendanceRecord model.AttendanceRecord) error
	EvaluteStudent(ctx context.Context, lesson_id int64, grade model.Grade) error
}

type People interface {
	Create(ctx context.Context, person model.CreatePersonInput) error
	GetAll(ctx context.Context) ([]model.Person, error)
	GetById(ctx context.Context, id int64) (model.Person, error)
	Update(ctx context.Context, id int64, person model.UpdatePersonInput) error
	Delete(ctx context.Context, id int64) error
	GetListApplications(ctx context.Context) ([]model.PersonApplication, error)
	GetApplication(ctx context.Context, personID int64) (model.PersonApplication, error)
	ResponseToApplication(ctx context.Context, response model.PersonApplication) error
	CreateApplicationPerson(ctx context.Context, person model.CreatePersonInput) error
}
type Admins interface {
	TryBlockUser(ctx context.Context, response model.ResponseComplaintInput) error
	UnblockUser(ctx context.Context, userID int64) error
}
type News interface {
	Create(ctx context.Context, news model.CreateNewsInput) error
	GetList(ctx context.Context) ([]model.News, error)
	GetNews(ctx context.Context, newsID int64) (model.News, error)
	Update(ctx context.Context, newsID int64, news model.UpdateNewsInput) error
	Delete(ctx context.Context, newsID int64) error
}
type Complaints interface {
	Create(ctx context.Context, complaint model.CreateComplaintInput) error
	GetAll(ctx context.Context) ([]model.Complaint, error)
	GetById(ctx context.Context, id int64) (model.Complaint, error)
	Response(ctx context.Context, complaintID int64, response model.ResponseComplaintInput) error
}
type Employees interface {
	Create(ctx context.Context, employee model.CreateEmployeeInput) error
	GetAll(ctx context.Context) ([]model.Employee, error)
	GetById(ctx context.Context, id int64) (model.Employee, error)
	Update(ctx context.Context, id int64, teacher model.UpdateEmployeeInput) error
	Delete(ctx context.Context, id int64) error
}
type Subjects interface {
	Create(ctx context.Context, subject model.CreateSubjectInput) error
	GetAll(ctx context.Context) ([]model.Subject, error)
	GetById(ctx context.Context, id int64) (model.Subject, error)
	Update(ctx context.Context, id int64, subject model.UpdateSubjectInput) error
	Delete(ctx context.Context, id int64) error
	GetStudentSubjects(ctx context.Context, student model.Student) ([]model.Subject, error)
}
type Lessons interface {
	Create(ctx context.Context, lesson model.CreateLessonInput) error
	GetAll(ctx context.Context) ([]model.Lesson, error)
	GetById(ctx context.Context, id int64) (model.Lesson, error)
	Delete(ctx context.Context, id int64) error
	StudentSchedule(ctx context.Context, student model.Student) ([]model.Lesson, error)
	TeacherSchedule(ctx context.Context, teacher model.Teacher) ([]model.Lesson, error)
}
type Faculties interface {
	Create(ctx context.Context, faculty model.CreateFacultyInput) error
	GetAll(ctx context.Context) ([]model.Faculty, error)
	GetById(ctx context.Context, id string) (model.Faculty, error)
	Delete(ctx context.Context, id string) error
}
type Specialties interface {
	Create(ctx context.Context, specialty model.CreateSpecialtyInput) error
	GetAll(ctx context.Context) ([]model.Specialty, error)
	GetById(ctx context.Context, id int64) (model.Specialty, error)
	Update(ctx context.Context, id int64, specialty model.UpdateSpecialtyInput) error
	Delete(ctx context.Context, id int64) error
	GetSpecialtiesByFacultyID(ctx context.Context, faculty_id string) ([]model.Specialty, error)
}
type Groups interface {
	Create(ctx context.Context, group model.CreateGroupInput) error
	GetAll(ctx context.Context) ([]model.Group, error)
	GetById(ctx context.Context, id string) (model.Group, error)
	Delete(ctx context.Context, id string) error
}
type Tokens struct {
	AccessToken  string
	RefreshToken string
}
