package psql

import (
	"context"
	"test-crud/internal/model"
	"time"

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
	Admins      Admins
	People      People
	Complaints  Complaints
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
		Admins:      NewAdminsRepository(db),
		People:      NewPeopleRepository(db),
		Complaints:  NewComplaintsRepository(db),
	}
}

type Students interface {
	Create(ctx context.Context, student model.CreateStudentInput) error
	GetAll(ctx context.Context) ([]model.Student, error)
	GetById(ctx context.Context, id int64) (model.Student, error)
	Update(ctx context.Context, id int64, student model.UpdateStudentInput) error
	Delete(ctx context.Context, id int64) error
	GetStudentBriefInfoById(ctx context.Context, id int64) (model.StudentBriefInfo, error)
	GetStudentFullInfoById(ctx context.Context, id int64) (model.StudentFullInfo, error)
}
type Users interface {
	Create(ctx context.Context, user model.User) error
	GetByEmailCredentials(ctx context.Context, login, password string) (model.User, error)
	GetByUsernameCredentials(ctx context.Context, login, password string) (model.User, error)
	SetSession(ctx context.Context, id int64, session model.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (model.User, error)
}
type Admins interface {
	BlockUser(ctx context.Context, id int64) error
	UnblockUser(ctx context.Context, id int64) error
}
type People interface {
	Create(ctx context.Context, person model.CreatePersonInput) error
	GetAll(ctx context.Context) ([]model.Person, error)
	GetById(ctx context.Context, id int64) (model.Person, error)
	Update(ctx context.Context, id int64, person model.UpdatePersonInput) error
	Delete(ctx context.Context, id int64) error
}
type Teachers interface {
	Create(ctx context.Context, teacher model.CreateTeacherInput) error
	GetAll(ctx context.Context) ([]model.Teacher, error)
	GetById(ctx context.Context, id int64) (model.Teacher, error)
	Update(ctx context.Context, id int64, teacher model.UpdateTeacherInput) error
	Delete(ctx context.Context, id int64) error
	GetTeacherBriefInfoById(ctx context.Context, id int64) (model.TeacherBriefInfo, error)
	GetTeacherFullInfoById(ctx context.Context, id int64) (model.TeacherFullInfo, error)
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
	GetSubjectsByStudentID(ctx context.Context, student_id int64) ([]model.Subject, error)
}
type Lessons interface {
	Create(ctx context.Context, lesson model.CreateLessonInput) error
	GetAll(ctx context.Context) ([]model.Lesson, error)
	GetById(ctx context.Context, id int64) (model.Lesson, error)
	Delete(ctx context.Context, id int64) error
	GetLessonsByStudentID(ctx context.Context, student_id int64, timeNow time.Time) ([]model.Lesson, error)
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
type Complaints interface {
	Create(ctx context.Context, complaint model.Complaint) error
	GetAll(ctx context.Context) ([]model.Complaint, error)
	GetById(ctx context.Context, id int64) (model.Complaint, error)
	Response(ctx context.Context, id int64, response string) error
}
