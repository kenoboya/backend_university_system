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
func (r *UsersRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (model.User, error) {
	var user model.User
	if err := r.db.Get(&user, "SELECT user_id, username, email, password, blocked,registered_at, last_visit_at, role FROM refresh_tokens JOIN users USING(user_id) WHERE token=$1 AND expires_at > NOW()", refreshToken); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UsersRepository) SetSession(ctx context.Context, userID uint64, session model.Session) error {
	_, err := r.db.Exec("DELETE FROM refresh_tokens WHERE user_id=$1", userID)
	if err != nil {
		return err
	}
	_, err1 := r.db.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES($1, $2, $3)", userID, session.RefreshToken, session.ExpiresAt)
	if err1 != nil {
		return err1
	}
	return nil
}

func (r *UsersRepository) UpdateRole(ctx context.Context, role string, user_id uint64) error {
	_, err := r.db.Exec("UPDATE users SET role=$1 WHERE user_id=$2", role, user_id)
	if err != nil {
		return err
	}
	return nil
}
