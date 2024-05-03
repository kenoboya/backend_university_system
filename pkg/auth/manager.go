package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken             = errors.New("invalid token")
	ErrorGettingClaimsFromToken = errors.New("error get claims from token")
	ErrEmptySecretKey           = errors.New("secret key is empty")
	ErrParseFloatToInt          = errors.New("id claim is not a valid float64")
)

type TokenManager interface {
	NewJWT(user_ID int64, role string, ttl time.Duration) (string, error)
	VerifyToken(accessToken string) (int64, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	secretKey string
}

func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, ErrEmptySecretKey
	}
	return &Manager{
		secretKey: secretKey,
	}, nil
}
func (m *Manager) NewJWT(user_ID int64, role string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user_ID,
		"role": role,
		"exp":  time.Now().Add(ttl).Unix(),
	})
	tokenString, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *Manager) VerifyToken(accessToken string) (int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, ErrorGettingClaimsFromToken
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, ErrParseFloatToInt
	}
	userID := int64(id)
	return userID, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
