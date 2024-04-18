package model

type Student struct {
	ID     int64 `db:"student_id" json:"student_id"`
	Person Person
}

type CreateStudentInput struct {
	PersonID int64 `db:"person_id" json:"person_id"`
	GroupID  int64 `db:"group_id" json:"group_id"`
}

// IN DEVELOPING
// type UpdateStudentInput struct {
// 	ID      int64  `db:"student_id" json:"student_id"`
// 	Person Person
// }
