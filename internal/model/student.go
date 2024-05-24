package model

type Student struct {
	StudentID uint64 `db:"student_id" json:"student_id"`
	PersonID  uint64 `db:"person_id" json:"person_id"`
	GroupID   string `db:"group_id" json:"group_id"`
}

type StudentBriefInfo struct {
	PersonBriefInfo
}
type StudentFullInfo struct {
	PersonFullInfo
}
type CreateStudentInput struct {
	PersonID uint64 `db:"person_id" json:"person_id"`
	GroupID  string `db:"group_id" json:"group_id"`
}

type UpdateStudentInput struct {
	GroupID string `db:"group_id" json:"group_id"`
}
