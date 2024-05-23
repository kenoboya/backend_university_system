package model

import (
	"time"
)

const (
	Blocked   = "blocked"
	Unblocked = "unblocked"

	RoleAdmin    = "admin"
	RoleTeacher  = "teacher"
	RoleStudent  = "student"
	RoleEmployee = "employee"
	RoleUser     = "user"
)

type User struct {
	UserID       int64      `db:"user_id" json:"user_id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	Password     string     `db:"password" json:"password"`
	RegisteredAt time.Time  `db:"registered_at" json:"registered_at"`
	LastVisitAt  *time.Time `db:"last_visit_at" json:"last_visit_at"`
	Blocked      string     `db:"blocked" json:"blocked,omitempty"`
	Role         string     `db:"role"`
	// Verification Verification `json:"verification"`
}

func (u *User) IsBlocked() bool {
	return u.Blocked == Blocked
}

type UserSignUpInput struct {
	Username string `json:"username" validate:"required,min=5,max=70"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type UserSignInInput struct {
	Login    string `json:"login" validate:"required,min=3"`
	Password string `json:"password" validate:"required,gte=6"`
	// Verified bool `json:"verified"`
	// Blocked bool `json:"blocked"`
}
