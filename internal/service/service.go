package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"test-crud/pkg/auth"
	"test-crud/pkg/hash"
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
}
type Deps struct {
	Repos        psql.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Students:    NewStudentsService(deps.Repos.Students),
		Users:       NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager),
		Teachers:    NewTeachersService(deps.Repos.Teachers),
		Employees:   NewEmployeesService(deps.Repos.Employees),
		Subjects:    NewSubjectsService(deps.Repos.Subjects),
		Lessons:     NewLessonsService(deps.Repos.Lessons),
		Faculties:   NewFacultiesService(deps.Repos.Faculties),
		Specialties: NewSpecialtiesService(deps.Repos.Specialties),
		Groups:      NewGroupsService(deps.Repos.Groups),
		Admins:      NewAdminsService(deps.Repos.Admins),
	}
}

type Students interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	GetAll(ctx context.Context) ([]model.Student, error)
	GetById(ctx context.Context, id int64) (model.Student, error)
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	Delete(ctx context.Context, id int64) error
}
type Users interface {
	SignUp(ctx context.Context, input model.UserSignUpInput) error
	SignIn(ctx context.Context, input model.UserSignInInput) (Tokens, error)
}
type Teachers interface {
	Create(ctx context.Context, teacher model.CreateTeacherInput) error
	GetAll(ctx context.Context) ([]model.Teacher, error)
	GetById(ctx context.Context, id int64) (model.Teacher, error)
	Update(ctx context.Context, id int64, teacher model.UpdateTeacherInput) error
	Delete(ctx context.Context, id int64) error
}

///

type Admins interface {
	// todo
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
}
type Lessons interface {
	Create(ctx context.Context, lesson model.CreateLessonInput) error
	GetAll(ctx context.Context) ([]model.Lesson, error)
	GetById(ctx context.Context, id int64) (model.Lesson, error)
	Delete(ctx context.Context, id int64) error
}
type Faculties interface {
	Create(ctx context.Context, faculty model.CreateFacultyInput) error
	GetAll(ctx context.Context) ([]model.Faculty, error)
	GetById(ctx context.Context, id int64) (model.Faculty, error)
	Delete(ctx context.Context, id int64) error
}
type Specialties interface {
	Create(ctx context.Context, specialty model.CreateSpecialtyInput) error
	GetAll(ctx context.Context) ([]model.Specialty, error)
	GetById(ctx context.Context, id int64) (model.Specialty, error)
	Update(ctx context.Context, id int64, specialty model.UpdateSpecialtyInput) error
	Delete(ctx context.Context, id int64) error
}
type Groups interface {
	Create(ctx context.Context, group model.CreateGroupInput) error
	GetAll(ctx context.Context) ([]model.Group, error)
	GetById(ctx context.Context, id int64) (model.Group, error)
	Delete(ctx context.Context, id int64) error
}
type Tokens struct {
	AccessToken  string
	RefreshToken string
}
