package model

type Teacher struct {
	TeacherID  uint64 `db:"teacher_id" json:"teacher_id"`
	EmployeeID uint64 `db:"employee_id" json:"employee_id"`
	// Subjects []Subject
}

type TeacherBriefInfo struct {
	PersonBriefInfo
}
type TeacherFullInfo struct {
	PersonFullInfo
	EmployeeBriefInfo
}

type CreateTeacherInput struct {
	EmployeeID uint64 `db:"employee_id" json:"employee_id"`
	// SEPARATE TABLE?
	// Subject_id []uint64 `db:"subject_id" json:"subject_id"`
}

type UpdateTeacherInput struct {
	EmployeeID uint64 `db:"employee_id" json:"employee_id"`
	// SEPARATE TABLE?
	// Subject_id []uint64 `db:"subject_id" json:"subject_id"`
}
