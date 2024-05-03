package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *Handler) initAdminLessonsRoutes(admin *mux.Router) {
	lessons := admin.PathPrefix("/lessons").Subrouter()
	{
		lessons.HandleFunc("", h.Admins.createLesson).Methods(http.MethodPost)
		lessons.HandleFunc("", h.Admins.getLessons).Methods(http.MethodGet)
		lessons.HandleFunc("/{id:[0-9]+}", h.Admins.getLesson).Methods(http.MethodGet)
		lessons.HandleFunc("/{id:[0-9]+}", h.Admins.deleteLesson).Methods(http.MethodDelete)
	}
}

// @Summary create lesson
// @Description create lesson
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param lesson body model.CreateLessonInput true "Data for creating lesson"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/lessons [post]
func (h *AdminsHandler) createLesson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "createLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var lesson model.CreateLessonInput
	if err := json.Unmarshal(reqBytes, &lesson); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "createLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Lessons.Create(context.TODO(), lesson); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "createLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get lessons
// @Description get lessons
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Success 200 {array} model.Lesson "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/lessons [get]
func (h *AdminsHandler) getLessons(w http.ResponseWriter, r *http.Request) {
	lessons, err := h.services.Lessons.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "getLessons()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lessons)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "getLessons()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get lesson
// @Description get lesson by id
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID for getting lesson"
// @Success 200 {object} model.Lesson "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/lessons/{id} [get]
func (h *AdminsHandler) getLesson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lesson, err := h.services.Lessons.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(lesson)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Delete lesson
// @Description delete lesson
// @Tags admin-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting lesson"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/lessons/{id} [delete]
func (h *AdminsHandler) deleteLesson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "deleteLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Lessons.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_lesson.go"),
			zap.String("function", "deleteLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
