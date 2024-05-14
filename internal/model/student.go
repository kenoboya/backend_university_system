package model

type Student struct {
	StudentID int64 `db:"student_id" json:"student_id"`
	Person
	Group
}

type StudentBriefInfo struct {
	StudentID int64 `db:"student_id" json:"student_id"`
	PersonBriefInfo
}
type StudentFullInfo struct {
	PersonFullInfo
}
type CreateStudentInput struct {
	PersonID int64 `db:"person_id" json:"person_id"`
	GroupID  int64 `db:"group_id" json:"group_id"`
}

type UpdateStudentInput struct {
	GroupID int64 `db:"group_id" json:"group_id"`
}
