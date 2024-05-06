package teacher

import (
	"test-crud/internal/service"
)

type TeachersHandler struct {
	studentService service.Students
	teacherService service.Teachers
}

func NewTeachersHandler(studentService service.Students, teacherService service.Teachers) *TeachersHandler {
	return &TeachersHandler{
		studentService: studentService,
		teacherService: teacherService,
	}
}
