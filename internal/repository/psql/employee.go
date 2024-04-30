package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type EmployeesRepository struct {
	db *sqlx.DB
}

func NewEmployeesRepository(db *sqlx.DB) *EmployeesRepository {
	return &EmployeesRepository{db}
}

func (r EmployeesRepository) Create(ctx context.Context, employee model.CreateEmployeeInput) error {
	_, err := r.db.NamedExec("INSERT INTO employees(person_id, title, salary, hire_date) VALUES(:person_id, :salary, :hire_date)", employee)
	if err != nil {
		return err
	}
	return nil
}
func (r EmployeesRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	employees := []model.Employee{}
	err := r.db.Select(&employees, "SELECT * FROM employees JOIN people USING(person_id)")
	if err != nil {
		return employees, err
	}
	return employees, nil
}
func (r EmployeesRepository) GetById(ctx context.Context, id int64) (model.Employee, error) {
	var employee model.Employee
	err := r.db.Get(&employee, "SELECT * FROM employees JOIN people USING(person_id) WHERE employee_id = $1", id)
	if err != nil {
		return employee, err
	}
	return employee, nil
}
func (r EmployeesRepository) Update(ctx context.Context, id int64, employee model.UpdateEmployeeInput) error {
	_, err := r.db.Exec("UPDATE employees SET title = $1, salary = $2, hire_date = $3 WHERE employee_id = $4",
		employee.Title, employee.Salary, employee.HireDate, id)
	if err != nil {
		return err
	}
	return nil
}
func (r EmployeesRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM employees WHERE employee_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
