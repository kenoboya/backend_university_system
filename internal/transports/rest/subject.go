package rest

import (
	"test-crud/internal/service"
)

type SubjectsHandler struct {
	service service.Subjects
}

func NewSubjectsHandler(service service.Subjects) *SubjectsHandler {
	return &SubjectsHandler{service: service}
}

type LessonsHandler struct {
	service service.Lessons
}

func NewLessonsHandler(service service.Lessons) *LessonsHandler {
	return &LessonsHandler{service: service}
}
