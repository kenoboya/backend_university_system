package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"test-crud/pkg/auth"
	"test-crud/pkg/hash"
)

type Services struct {
	//Students    Students
	Users       Users
	Teachers    Teachers
	Employees   Employees
	Subjects    Subjects
	Lessons     Lessons
	Faculties   Faculties
	Specialties Specialties
	Groups      Groups
}
type Deps struct {
	Repos        psql.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		//Students:    NewStudentsService(deps.Repos.Students),
		Users:       NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager),
		Teachers:    NewTeachersService(deps.Repos.Teachers),
		Employees:   NewEmployeesService(deps.Repos.Employees),
		Subjects:    NewSubjectsService(deps.Repos.Subjects),
		Lessons:     NewLessonsService(deps.Repos.Lessons),
		Faculties:   NewFacultiesService(deps.Repos.Faculties),
		Specialties: NewSpecialtiesService(deps.Repos.Specialties),
		Groups:      NewGroupsService(deps.Repos.Groups),
	}
}

type Students interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	GetById(ctx context.Context, id int64) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
}
type Users interface {
	SignUp(ctx context.Context, input model.UserSignUpInput) error
	SignIn(ctx context.Context, input model.UserSignInInput) (Tokens, error)
}
type Teachers interface {
	// todo
}
type Employees interface {
	// todo
}
type Subjects interface {
	// todo
}
type Lessons interface {
	// todo
}
type Specialties interface {
	// todo
}
type Faculties interface {
	// todo
}
type Groups interface {
	// todo
}
type Tokens struct {
	AccessToken  string
	RefreshToken string
}
