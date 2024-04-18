package model

import "time"

type User struct {
	ID           int64     `db:"person_id" json:"person_id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	Password     string    `db:"password" json:"password"`
	RegisteredAt time.Time `db:"registered_at" json:"registered_at"`
	LastVisitAt  time.Time `db:"last_visit_at" json:"last_visit_at"`
	// Verification Verification `json:"verification"`
	// Session Session json:"session"`
	// Blocked bool json:"blocked"`
}

type CreateUserInput struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UpdateUserInput struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	// Verified bool `json:"verified"`
	// Blocked bool `json:"blocked"`
}
