package model

type Teacher struct {
	ID       int64 `db:"teacher_id" json:"teacher_id"`
	Employee Employee
	// Subjects []Subject
}
type CreateTeacherInput struct {
	EmployeeID int64 `db:"employee_id" json:"employee_id"`
	// SEPARATE TABLE?
	// Subject_id []int64 `db:"subject_id" json:"subject_id"`
}

// IN DEVELOPING
// type UpdateTeacherInput struct {
// 	SEPARATE TABLE?
// 	Subject_id []int64 `db:"subject_id" json:"subject_id"`
// }
