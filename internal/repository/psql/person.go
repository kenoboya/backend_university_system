package psql

import (
	"context"
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
	_, err := r.db.NamedExec("INSERT INTO people(user_id,name,surname,birth_date,phone) VALUES(:user_id,:name,:surname,:birth_date, phone)", person)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) GetAll(ctx context.Context) ([]model.Person, error) {
	people := []model.Person{}
	err := r.db.Select(&people, "SELECT * FROM people JOIN users USING(user_id)")
	if err != nil {
		return people, err
	}
	return people, err
}

func (r *PeopleRepository) GetById(ctx context.Context, id int64) (model.Person, error) {
	var person model.Person
	err := r.db.Get(&person, "SELECT * FROM people JOIN users USING(user_id) WHERE person_id=$1", id)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *PeopleRepository) GetPersonByUserID(ctx context.Context, userID int64) (model.Person, error) {
	var person model.Person
	err := r.db.Get(&person, "SELECT * FROM people JOIN users USING(user_id) WHERE user_id=$1", userID)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (r *PeopleRepository) Update(ctx context.Context, id int64, person model.UpdatePersonInput) error {
	_, err := r.db.Exec("UPDATE people SET name=$1,surname=$2,birth_date=$3,phone=$4,address=$5,photo=$6,notes=$7 WHERE person_id=$8",
		person.Name, person.Surname, person.Birth_date, person.Phone, person.Address, person.Photo, person.Notes, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM people WHERE person_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) GetAllApplications(ctx context.Context) ([]model.PersonApplication, error) {
	applications := []model.PersonApplication{}
	err := r.db.Select(&applications, "SELECT * FROM applications_people JOIN people USING(person_id) JOIN users USING(user_id)")
	if err != nil {
		return applications, err
	}
	return applications, err
}
func (r *PeopleRepository) GetApplication(ctx context.Context, personID int64) (model.PersonApplication, error) {
	var application model.PersonApplication
	err := r.db.Get(&application, "SELECT * FROM applications_people JOIN people USING(person_id) JOIN users USING(user_id) WHERE person_id=$1", personID)
	if err != nil {
		return application, err
	}
	return application, nil
}
func (r *PeopleRepository) UpdateApplicationStatus(ctx context.Context, response bool, id int64) error {
	_, err := r.db.Exec("UPDATE applications_people SET accepted=$1, WHERE application_id=$2", response, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PeopleRepository) CreateApplicationPerson(ctx context.Context, personID int64) error {
	_, err := r.db.NamedExec("INSERT INTO applications_people(person_id) VALUES(:person_id)", personID)
	if err != nil {
		return err
	}
	return nil
}
