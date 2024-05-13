package student

import (
	"test-crud/internal/service"
)

type StudentsHandler struct {
	studentService service.Students
	teacherService service.Teachers
	subjectService service.Subjects
	lessonService  service.Lessons
}

func NewStudentsHandler(studentService service.Students, teacherService service.Teachers, subjectService service.Subjects, lessonService service.Lessons) *StudentsHandler {
	return &StudentsHandler{
		studentService: studentService,
		teacherService: teacherService,
		subjectService: subjectService,
		lessonService:  lessonService,
	}
}
