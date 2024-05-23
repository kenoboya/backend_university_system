package psql

import (
	"context"
	"fmt"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type PeopleRepository struct {
	db *sqlx.DB
}

func NewPeopleRepository(db *sqlx.DB) *PeopleRepository {
	return &PeopleRepository{db}
}

func (r *PeopleRepository) Create(ctx context.Context, person model.CreatePersonInput) error {
	_, err := r.db.NamedExec("INSERT INTO people(user_id,name,surname,birth_date,phone) VALUES(:user_id,:name,:surname,:birth_date, :phone)", person)
	if err != nil {
		fmt.Println("\nError: ", err)
		fmt.Println(person.Phone)
		return err
	}
	return nil
}

func (r *PeopleRepository) GetAll(ctx context.Context) ([]model.Person, error) {
	people := []model.Person{}
	err := r.db.Select(&people, "SELECT * FROM people")
	if err != nil {
		return people, err
	}
	return people, err
}

func (r *PeopleRepository) GetById(ctx context.Context, id uint64) (model.Person, error) {
	var person model.Person
	err := r.db.Get(&person, "SELECT * FROM people WHERE person_id=$1", id)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *PeopleRepository) GetPersonByUserID(ctx context.Context, userID uint64) (model.Person, error) {
	var person model.Person
	err := r.db.Get(&person, "SELECT * FROM people WHERE user_id=$1", userID)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *PeopleRepository) Update(ctx context.Context, id uint64, person model.UpdatePersonInput) error {
	_, err := r.db.Exec("UPDATE people SET name=$1,surname=$2,birth_date=$3,phone=$4,address=$5,photo=$6,notes=$7 WHERE person_id=$8",
		person.Name, person.Surname, person.Birth_date, person.Phone, person.Address, person.Photo, person.Notes, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) Delete(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("DELETE FROM people WHERE person_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) GetAllApplications(ctx context.Context) ([]model.PersonApplication, error) {
	applications := []model.PersonApplication{}
	err := r.db.Select(&applications, "SELECT * FROM applications_people")
	if err != nil {
		return applications, err
	}
	return applications, err
}
func (r *PeopleRepository) GetApplicationByID(ctx context.Context, applicationID uint64) (model.PersonApplication, error) {
	var application model.PersonApplication
	err := r.db.Get(&application, "SELECT * FROM applications_people WHERE application_id=$1", applicationID)
	if err != nil {
		return application, err
	}
	return application, nil
}
func (r *PeopleRepository) GetApplicationsByUserID(ctx context.Context, userID uint64) ([]model.PersonApplication, error) {
	applications := []model.PersonApplication{}
	err := r.db.Select(&applications, "SELECT * FROM applications_people WHERE user_id=$1", userID)
	if err != nil {
		return applications, err
	}
	return applications, err
}
func (r *PeopleRepository) UpdateApplicationStatus(ctx context.Context, status string, id uint64) error {
	_, err := r.db.Exec("UPDATE applications_people SET status=$1 WHERE application_id=$2", status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) CreateApplicationPerson(ctx context.Context, input model.CreatePersonInput) error {
	_, err := r.db.Exec("INSERT INTO applications_people(user_id, role) VALUES($1, $2)", input.UserID, input.Role)
	if err != nil {
		return err
	}
	return nil
}
