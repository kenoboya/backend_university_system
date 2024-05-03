package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type ComplaintsRepository struct {
	db *sqlx.DB
}

func NewComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db}
}

func (r *ComplaintsRepository) Create(ctx context.Context, complaint model.Complaint) error {
	_, err := r.db.NamedExec("INSERT INTO complaints (reporting_user_id, reported_user_id, cause, time, answer) VALUES(:reporting_user_id, :reported_user_id, :cause, :time, :answer)", complaint)
	if err != nil {
		return err
	}
	return nil
}

func (r *ComplaintsRepository) GetAll(ctx context.Context) ([]model.Complaint, error) {
	complaints := []model.Complaint{}
	err := r.db.Select(&complaints, "SELECT * FROM complaints")
	if err != nil {
		return complaints, err
	}
	return complaints, nil
}

func (r *ComplaintsRepository) GetById(ctx context.Context, id int64) (model.Complaint, error) {
	complaint := model.Complaint{}
	err := r.db.Get(&complaint, "SELECT * FROM complaints WHERE complaint_id=$1", id)
	if err != nil {
		return complaint, err
	}
	return complaint, nil
}

func (r *ComplaintsRepository) Response(ctx context.Context, id int64, response string) error {
	_, err := r.db.Exec("UPDATE complaints SET response=$1 WHERE complaint_id=$2", response, id)
	if err != nil {
		return err
	}
	return nil
}
