package model

import "time"

const (
	Accepted      = "accepted"
	Consideration = "consideration"
	Denied        = "denied"
)

type Person struct {
	PersonID  uint64    `db:"person_id" json:"person_id"`
	UserID    uint64    `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Phone     string    `db:"phone" json:"phone"`
	Address   *string   `db:"address" json:"address"`
	Photo     *[]byte   `db:"photo" json:"photo"`
	Notes     *string   `db:"notes" json:"notes"`
}

type PersonBriefInfo struct {
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Photo     *[]byte   `db:"photo" json:"photo"`
}

type PersonFullInfo struct {
	PersonBriefInfo
	Phone   string  `db:"phone" json:"phone"`
	Address *string `db:"address" json:"address"`
}

type CreatePersonInput struct {
	UserID    uint64    `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Phone     string    `db:"phone" json:"phone"`
	Role      string    `json:"role"`
}
type UpdatePersonInput struct {
	Name       string    `db:"name" json:"name"`
	Surname    string    `db:"surname" json:"surname"`
	Birth_date time.Time `db:"birth_date" json:"birth_date"`
	Phone      string    `db:"phone" json:"phone"`
	Address    *string   `db:"address" json:"address"`
	Photo      *[]byte   `db:"photo" json:"photo"`
	Notes      *string   `db:"notes" json:"notes"`
}

type PersonApplication struct {
	ApplicationID uint64 `db:"application_id" json:"application_id"`
	UserID        uint64 `db:"user_id" json:"user_id"`
	Role          string `db:"role" json:"role"`
	Status        string `db:"status" json:"status"`
}
