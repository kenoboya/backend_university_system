package model

type Student struct {
	StudentID uint64 `db:"student_id" json:"student_id"`
	Person
	Group
}

type StudentBriefInfo struct {
	StudentID uint64 `db:"student_id" json:"student_id"`
	PersonBriefInfo
}
type StudentFullInfo struct {
	PersonFullInfo
}
type CreateStudentInput struct {
	PersonID uint64 `db:"person_id" json:"person_id"`
	GroupID  uint64 `db:"group_id" json:"group_id"`
}

type UpdateStudentInput struct {
	GroupID uint64 `db:"group_id" json:"group_id"`
}
