package model

import "time"

const (
	LessonLecture    = "lecture"
	LessonLaboratory = "laboratory"
	LessonPractical  = "Practical"
	LessonSession    = "Session"
)

type Faculty struct {
	FacultyID   string `db:"faculty_id" json:"faculty_id"`
	FacultyName string `db:"full_name" json:"full_name"`
}
type Specialty struct {
	SpecialtyID   int64  `db:"specialty_id" json:"specialty_id"`
	SpecialtyName string `db:"full_name" json:"full_name"`
	Faculty
}
type Group struct {
	GroupID          string    `db:"group_id" json:"group_id"`
	GroupName        string    `db:"full_name" json:"full_name"`
	EducationalLevel string    `db:"educational_level" json:"educational_level"`
	StartYear        time.Time `db:"start_year" json:"start_year"`
	EndYear          time.Time `db:"end_year" json:"end_year"`
	Specialty
}

type CreateFacultyInput struct {
	FullName string `db:"full_name" json:"full_name"`
}

type CreateSpecialtyInput struct {
	SpecialtyID int64  `db:"specialty_id" json:"specialty_id"`
	FacultyID   string `db:"faculty_id" json:"faculty_id"`
	FullName    string `db:"full_name" json:"full_name"`
}
type UpdateSpecialtyInput struct {
	FacultyID string `db:"faculty_id" json:"faculty_id"`
	FullName  string `db:"full_name" json:"full_name"`
}

type CreateGroupInput struct {
	SpecialtyID      string    `db:"specialty_id" json:"specialty_id"`
	FullName         string    `db:"full_name" json:"full_name"`
	StartYear        time.Time `db:"start_year" json:"start_year"`
	EducationalLevel string    `db:"educational_level" json:"educational_level"`
	// Most likely there will be different degrees of education (types or enums): bachelor, master, doctor,
	// From their will be depend EndYear -> processing will be in the service
	// EndYear   time.Time `db:"end_year" json:"end_year"`
}
