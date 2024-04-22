package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type User struct {
	ID           int64      `db:"person_id" json:"person_id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	Password     string     `db:"password" json:"password"`
	RegisteredAt time.Time  `db:"registered_at" json:"registered_at"`
	LastVisitAt  *time.Time `db:"last_visit_at" json:"last_visit_at"`
	// Verification Verification `json:"verification"`
	// Session Session json:"session"`
	// Blocked bool json:"blocked"`
}

type UserSignUpInput struct {
	Username string `json:"username" validate:"required, min=5, max=70"`
	Email    string `json:"email" validate:"required, min=3"`
	Password string `json:"password" validate:"required, gte=6"`
}

func (i UserSignUpInput) Validator() error {
	return validate.Struct(i)
}

type UserSignInInput struct {
	Login    string `json:"login" validate:"required, min=3"`
	Password string `json:"password" validate:"required, gte=6"`
	// Verified bool `json:"verified"`
	// Blocked bool `json:"blocked"`
}

func (i UserSignInInput) Validator() error {
	return validate.Struct(i)
}
