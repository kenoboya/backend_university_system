package rest

import (
	"test-crud/internal/service"

	"github.com/gorilla/mux"
)

type SubjectsHandler struct {
	service service.Subjects
}

func NewSubjectsHandler(service service.Subjects) *SubjectsHandler {
	return &SubjectsHandler{service: service}
}

func (h *SubjectsHandler) initRoutes(router *mux.Router) {
	// subjects := router.PathPrefix("/subjects").Subrouter()
	// {
	// 	// todo
	// }
}

type LessonsHandler struct {
	service service.Lessons
}

func NewLessonsHandler(service service.Lessons) *LessonsHandler {
	return &LessonsHandler{service: service}
}

func (h *LessonsHandler) initRoutes(router *mux.Router) {
	// lessons := router.PathPrefix("/lessons").Subrouter()
	// {
	// 	// todo
	// }
}
