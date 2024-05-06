package model

type Teacher struct {
	TeacherID int64 `db:"teacher_id" json:"teacher_id"`
	Employee
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
	EmployeeID int64 `db:"employee_id" json:"employee_id"`
	// SEPARATE TABLE?
	// Subject_id []int64 `db:"subject_id" json:"subject_id"`
}

type UpdateTeacherInput struct {
	EmployeeID int64 `db:"employee_id" json:"employee_id"`
	// SEPARATE TABLE?
	// Subject_id []int64 `db:"subject_id" json:"subject_id"`
}
