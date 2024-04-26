package model

import "time"

type Faculty struct {
	ID       string `db:"faculty_id" json:"faculty_id"`
	FullName string `db:"full_name" json:"full_name"`
}
type Specialty struct {
	ID       int64  `db:"specialty_id" json:"specialty_id"`
	FullName string `db:"full_name" json:"full_name"`
	Faculty  Faculty
}
type Group struct {
	ID               string    `db:"specialty_id" json:"specialty_id"`
	FullName         string    `db:"full_name" json:"full_name"`
	EducationalLevel string    `db:"educational_level" json:"educational_level"`
	StartYear        time.Time `db:"start_year" json:"start_year"`
	EndYear          time.Time `db:"end_year" json:"end_year"`
	Specialty        Specialty
}

type CreateFacultyInput struct {
	FullName string `db:"full_name" json:"full_name"`
}

type CreateSpecialtyInput struct {
	FacultyID string `db:"faculty_id" json:"faculty_id"`
	FullName  string `db:"full_name" json:"full_name"`
}
type UpdateSpecialtyInput struct {
	FacultyID string `db:"faculty_id" json:"faculty_id"`
	FullName  string `db:"full_name" json:"full_name"`
}

type CreateGroupInput struct {
	SpecialtyID string    `db:"specialty_id" json:"specialty_id"`
	FullName    string    `db:"full_name" json:"full_name"`
	StartYear   time.Time `db:"start_year" json:"start_year"`
	// Most likely there will be different degrees of education (types or enums): bachelor, master, doctor,
	// From their will be depend EndYear -> processing will be in the service
	// EndYear   time.Time `db:"end_year" json:"end_year"`
}
type UpdateGroupInput struct {
	FullName  string    `db:"full_name" json:"full_name"`
	StartYear time.Time `db:"start_year" json:"start_year"`
	// Most likely there will be different degrees of education (types or enums): bachelor, master, doctor,
	// From their will be depend EndYear -> processing will be in the service
	// EndYear   time.Time `db:"end_year" json:"end_year"`
}
