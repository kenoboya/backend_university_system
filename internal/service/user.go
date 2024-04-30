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

func NewUsersService(repo psql.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		repo:         repo,
		hasher:       hasher,
		tokenManager: tokenManager,
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

func (s *UsersService) Refresh(ctx context.Context, refreshToken string) (Tokens, error) {
	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(ctx, user.ID)
}
func (s *UsersService) createSession(ctx context.Context, userID int64) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken, err = s.tokenManager.NewJWT(userID, s.accessTokenTTL)
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
	err = s.repo.SetSession(ctx, userID, session)
	return res, err
}
