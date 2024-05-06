package student

import (
	"test-crud/internal/service"
)

type StudentsHandler struct {
	studentService service.Students
	teacherService service.Teachers
}

func NewStudentsHandler(studentService service.Students, teacherService service.Teachers) *StudentsHandler {
	return &StudentsHandler{
		studentService: studentService,
		teacherService: teacherService,
	}
}
