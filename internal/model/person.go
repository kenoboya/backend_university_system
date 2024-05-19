package model

import "time"

type Person struct {
	PersonID  int64     `db:"person_id" json:"person_id"`
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Phone     string    `db:"phone" json:"phone"`
	Address   string    `db:"address" json:"address"`
	Photo     []byte    `db:"photo" json:"photo"`
	Notes     string    `db:"notes" json:"notes"`
	User
}

type PersonBriefInfo struct {
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Photo     []byte    `db:"photo" json:"photo"`
}

type PersonFullInfo struct {
	PersonBriefInfo
	Phone   string `db:"phone" json:"phone"`
	Address string `db:"address" json:"address"`
}

type CreatePersonInput struct {
	UserID    int64     `db:"user_id" json:"user_id"`
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
	ApplicationID int64 `db:"application_id" json:"application_id"`
	Person
	Accepted bool `db:"accepted" json:"accepted"`
}
