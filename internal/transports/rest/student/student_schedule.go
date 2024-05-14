package student

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *StudentsHandler) InitStudentScheduleRoutes(hub *mux.Router) {
	schedule := hub.PathPrefix("/schedule").Subrouter()
	{
		schedule.HandleFunc("", h.Schedule).Methods(http.MethodGet)
		// schedule.HandleFunc("/{id:[0-9]+}", h.GetStudentSubject).Methods(http.MethodGet)
	}
}

// @Summary Get student's schedule
// @Description get student's schedule
// @Tags student-schedule
// @Accept json
// @Produce json
// @Param student body model.Student true "student"
// @Success 200 {array} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /student/hub/schedule [get]
func (h *StudentsHandler) Schedule(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_schedule.go"),
			zap.String("function", "StudentSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student model.Student

	if err = json.Unmarshal(reqBytes, &student); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_schedule.go"),
			zap.String("function", "StudentSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lessons, err := h.lessonService.StudentSchedule(context.TODO(), student)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_schedule.go"),
			zap.String("function", "StudentSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lessons)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_schedule.go"),
			zap.String("function", "StudentSchedule"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
