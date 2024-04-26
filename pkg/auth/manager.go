package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	NewJWT(user_ID int64, ttl time.Duration) (string, error)
	VerifyToken(accessToken string) error
	NewRefreshToken() (string, error)
}

type Manager struct {
	secretKey string
}

func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("Secret key is empty")
	}
	return &Manager{
		secretKey: secretKey,
	}, nil
}
func (m *Manager) NewJWT(user_ID int64, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user_ID,
		"expiresAt": time.Now().Add(ttl).Unix(),
	})
	tokenString, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(accessToken string) error {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Invalid token")
	}
	return nil
}
