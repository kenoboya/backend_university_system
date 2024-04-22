package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"time"
)

type UsersService struct {
	repo psql.Users
	//hasher     hash.PasswordHasher
	//hmacSecret []byte
	//tokenTtl   time.Duration
}

func NewUsersService(repo psql.Users,

// hasher hash.PasswordHasher, secret []byte, ttl time.Duration
) *UsersService {
	return &UsersService{
		repo: repo,
		// hasher:     hasher,
		// hmacSecret: secret,
		// tokenTtl:   ttl,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input model.UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	user := model.User{
		Username:     input.Username,
		Password:     passwordHash,
		Email:        input.Email,
		RegisteredAt: time.Now(),
	}
	return s.repo.Create(ctx, user)
}

func (s *UsersService) SignIn(ctx context.Context, input model.UserSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}
	user, err := s.repo.GetByEmailCredentials(ctx, input.Login, passwordHash)
	if err != nil {
		user, err := s.repo.GetByUsernameCredentials(ctx, input.Login, passwordHash)
		if err != nil {
			return Tokens{}, err
		}
		s.createSession(ctx, user.ID)
	}
	return s.createSession(ctx, user.ID)
}
func (s *UsersService) createSession(ctx context.Context, userID int64) (Tokens, error) {
	// todo
}
