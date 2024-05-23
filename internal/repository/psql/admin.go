package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type AdminsRepository struct {
	db *sqlx.DB
}

func NewAdminsRepository(db *sqlx.DB) *AdminsRepository {
	return &AdminsRepository{db}
}

func (r *AdminsRepository) BlockUser(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("UPDATE users SET blocked=$1 WHERE user_id=$2", model.Blocked, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdminsRepository) UnblockUser(ctx context.Context, id uint64) error {
	_, err := r.db.Exec("UPDATE users SET blocked=$1 WHERE user_id=$2", model.Unblocked, id)
	if err != nil {
		return err
	}
	return nil
}
