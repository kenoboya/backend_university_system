package teacher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *TeachersHandler) InitTeacherScheduleRoutes(hub *mux.Router) {
	schedule := hub.PathPrefix("/schedule").Subrouter()
	{
		schedule.HandleFunc("", h.Schedule).Methods(http.MethodGet)
	}
}

// @Summary Get teacher's schedule
// @Description get teacher's schedule
// @Tags teacher-schedule
// @Accept json
// @Produce json
// @Param teacher body model.Teacher true "Teacher"
// @Success 200 {array} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/schedule [get]
func (h *TeachersHandler) Schedule(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_schedule.go"),
			zap.String("function", "TeacherSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var teacher model.Teacher

	if err = json.Unmarshal(reqBytes, &teacher); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_schedule.go"),
			zap.String("function", "TeacherSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lessons, err := h.lessonService.TeacherSchedule(context.TODO(), teacher)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_schedule.go"),
			zap.String("function", "TeacherSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lessons)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_schedule.go"),
			zap.String("function", "TeacherSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
