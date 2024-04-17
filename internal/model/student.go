package model

import "time"

type Student struct {
	ID           int64      `db:"student_id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Surname      string     `db:"surname" json:"surname"`
	Age          uint8      `db:"age" json:"age"`
	Email        string     `db:"email" json:"email"`
	Phone        string     `db:"phone" json:"phone"`
	RegisteredAt time.Time  `db:"registered_at" json:"registered_at"`
	LastVisitAt  *time.Time `db:"last_visit_at" json:"last_visit_at"`
}
type StudentBriefInfo struct {
	Name    string `db:"name" json:"name"`
	Surname string `db:"surname" json:"surname"`
	Age     uint8  `db:"age" json:"age"`
}
type UpdateStudentInput struct {
	ID      int64  `db:"student_id" json:"id"`
	Name    string `db:"name" json:"name"`
	Surname string `db:"surname" json:"surname"`
	Age     uint8  `db:"age" json:"age"`
	Email   string `db:"email" json:"email"`
	Phone   string `db:"phone" json:"phone"`
}
type CreateStudentInput struct {
	Name         string    `db:"name" json:"name"`
	Surname      string    `db:"surname" json:"surname"`
	Age          uint8     `db:"age" json:"age"`
	Email        string    `db:"email" json:"email"`
	Phone        string    `db:"phone" json:"phone"`
	RegisteredAt time.Time `db:"registered_at" json:"registered_at"`
}
