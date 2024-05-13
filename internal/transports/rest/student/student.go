package student

import (
	"test-crud/internal/service"
)

type StudentsHandler struct {
	studentService service.Students
	teacherService service.Teachers
	subjectService service.Subjects
}

func NewStudentsHandler(studentService service.Students, teacherService service.Teachers, subjectService service.Subjects) *StudentsHandler {
	return &StudentsHandler{
		studentService: studentService,
		teacherService: teacherService,
		subjectService: subjectService,
	}
}
