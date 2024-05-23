package model

import "time"

type Employee struct {
	EmployeeID uint64    `db:"employee_id" json:"employee_id"`
	PersonID   uint64    `db:"person_id" json:"person_id"`
	Title      string    `db:"title" json:"title"`
	Salary     float64   `db:"salary" json:"salary"`
	HireDate   time.Time `db:"hire_date" json:"hire_date"`
}

type EmployeeBriefInfo struct {
	Title    string    `db:"title" json:"title"`
	HireDate time.Time `db:"hire_date" json:"hire_date"`
}

type CreateEmployeeInput struct {
	PersonID uint64    `db:"person_id" json:"person_id"`
	Title    string    `db:"title" json:"title"`
	Salary   float64   `db:"salary" json:"salary"`
	HireDate time.Time `db:"hire_date" json:"hire_date"`
}
type UpdateEmployeeInput struct {
	Title    string    `db:"title" json:"title"`
	Salary   float64   `db:"salary" json:"salary"`
	HireDate time.Time `db:"hire_date" json:"hire_date"`
}
