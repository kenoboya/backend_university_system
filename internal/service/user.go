package service

import (
	"context"
	"test-crud/internal/model"
	"test-crud/internal/repository/psql"
	"test-crud/pkg/auth"
	"test-crud/pkg/hash"
	"time"
)

type UsersService struct {
	repo         psql.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(repo psql.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *UsersService {
	return &UsersService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
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
		Blocked:      model.Unblocked,
		Role:         model.RoleUser,
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
		if user.IsBlocked() {
			return Tokens{}, model.ErrUserBlocked
		}
		return s.createSession(ctx, &user)
	}
	if user.IsBlocked() {
		return Tokens{}, model.ErrUserBlocked
	}
	return s.createSession(ctx, &user)
}

func (s *UsersService) Refresh(ctx context.Context, refreshToken string) (Tokens, error) {
	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if user.IsBlocked() {
		return Tokens{}, err
	}
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(ctx, &user)
}
func (s *UsersService) createSession(ctx context.Context, user *model.User) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken, err = s.tokenManager.NewJWT(user.UserID, user.Role, s.accessTokenTTL)
	if err != nil {
		return res, err
	}
	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := model.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}
	err = s.repo.SetSession(ctx, user.UserID, session)
	return res, err
}
func (s *UsersService) ChangeRole(ctx context.Context, role string, user_id uint64) error {
	return s.repo.UpdateRole(ctx, role, user_id)
}
