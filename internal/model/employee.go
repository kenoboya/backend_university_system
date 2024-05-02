package model

import "time"

type Employee struct {
	EmployeeID int64     `db:"employee_id" json:"employee_id"`
	Title      string    `db:"title" json:"title"`
	Salary     float64   `db:"salary" json:"salary"`
	HireDate   time.Time `db:"hire_date" json:"hire_date"`
	Person
}

type CreateEmployeeInput struct {
	PersonID int64     `db:"person_id" json:"person_id"`
	Title    string    `db:"title" json:"title"`
	Salary   float64   `db:"salary" json:"salary"`
	HireDate time.Time `db:"hire_date" json:"hire_date"`
}
type UpdateEmployeeInput struct {
	Title    string    `db:"title" json:"title"`
	Salary   float64   `db:"salary" json:"salary"`
	HireDate time.Time `db:"hire_date" json:"hire_date"`
}
