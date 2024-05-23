package model

import "time"

const (
	Bachelor = "bachelor"
	Master   = "Master"
	Doctor   = "Doctor"
)

type Faculty struct {
	FacultyID   string `db:"faculty_id" json:"faculty_id"`
	FacultyName string `db:"full_name" json:"full_name"`
}
type Specialty struct {
	SpecialtyID   uint16 `db:"specialty_id" json:"specialty_id"`
	FacultyID     string `db:"faculty_id" json:"faculty_id"`
	SpecialtyName string `db:"full_name" json:"full_name"`
}
type Group struct {
	GroupID          string    `db:"group_id" json:"group_id"`
	SpecialtyID      uint16    `db:"specialty_id" json:"specialty_id"`
	GroupName        string    `db:"full_name" json:"full_name"`
	EducationalLevel string    `db:"educational_level" json:"educational_level"`
	StartYear        time.Time `db:"start_year" json:"start_year"`
	EndYear          time.Time `db:"end_year" json:"end_year"`
}

type CreateFacultyInput struct {
	FacultyID string `db:"faculty_id" json:"faculty_id"`
	FullName  string `db:"full_name" json:"full_name"`
}

type CreateSpecialtyInput struct {
	SpecialtyID uint16 `db:"specialty_id" json:"specialty_id"`
	FacultyID   string `db:"faculty_id" json:"faculty_id"`
	FullName    string `db:"full_name" json:"full_name"`
}
type UpdateSpecialtyInput struct {
	FacultyID string `db:"faculty_id" json:"faculty_id"`
	FullName  string `db:"full_name" json:"full_name"`
}

type CreateGroupInput struct {
	GroupID          string    `db:"group_id" json:"group_id"`
	SpecialtyID      uint16    `db:"specialty_id" json:"specialty_id"`
	FullName         string    `db:"full_name" json:"full_name"`
	StartYear        time.Time `db:"start_year" json:"start_year"`
	EducationalLevel string    `db:"educational_level" json:"educational_level"`
	// Most likely there will be different degrees of education (types or enums): bachelor, master, doctor,
	// From their will be depend EndYear -> processing will be in the service
	// EndYear   time.Time `db:"end_year" json:"end_year"`
}
