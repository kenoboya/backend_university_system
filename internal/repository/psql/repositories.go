package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
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

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Students:    NewStudentsRepository(db),
		Users:       NewUsersRepository(db),
		Teachers:    NewTeachersRepository(db),
		Employees:   NewEmployeesRepository(db),
		Subjects:    NewSubjectsRepository(db),
		Lessons:     NewLessonsRepository(db),
		Faculties:   NewFacultiesRepository(db),
		Specialties: NewSpecialtiesRepository(db),
		Groups:      NewGroupsRepository(db),
	}
}

// AdminService
type Students interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	GetById(ctx context.Context, id int64) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
}
type Users interface {
	Create(ctx context.Context, user model.User) error
	GetByEmailCredentials(ctx context.Context, login, password string) (model.User, error)
	GetByUsernameCredentials(ctx context.Context, login, password string) (model.User, error)
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
type Groups interface {
	// todo
}
type Faculties interface {
	// todo
}
