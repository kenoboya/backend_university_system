package psql

import (
	"context"
	"test-crud/internal/model"

	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) Create(ctx context.Context, user model.User) error {
	_, err := r.db.NamedExec("INSERT INTO users(username, email, password, registered_at) VALUES(:username, :email, :password, :registered_at)", &user)
	if err != nil {
		return err
	}
	return nil
}
func (r *UsersRepository) GetByEmailCredentials(ctx context.Context, login, password string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1 AND password = $2", login, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *UsersRepository) GetByUsernameCredentials(ctx context.Context, username, password string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1 AND password = $2", username, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
