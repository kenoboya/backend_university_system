package teacher

import (
	"test-crud/internal/service"
)

type TeachersHandler struct {
	studentService service.Students
	teacherService service.Teachers
	lessonService  service.Lessons
}

func NewTeachersHandler(studentService service.Students, teacherService service.Teachers, lessonService service.Lessons) *TeachersHandler {
	return &TeachersHandler{
		studentService: studentService,
		teacherService: teacherService,
		lessonService:  lessonService,
	}
}
