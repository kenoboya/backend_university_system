package student

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"
	"test-crud/internal/transports/rest/common"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *StudentsHandler) InitStudentSubjectsRoutes(hub *mux.Router) {
	subjects := hub.PathPrefix("/subjects").Subrouter()
	{
		subjects.HandleFunc("", h.GetStudentSubjects).Methods(http.MethodGet)
		subjects.HandleFunc("/{id:[0-9]+}", h.GetStudentSubject).Methods(http.MethodGet)
	}
}

// @Summary Get student's subjects
// @Description get student's subjects
// @Tags student-subjects
// @Accept json
// @Produce json
// @Param student body model.Student true "student"
// @Success 200 {array} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /student/hub/subjects [get]
func (h *StudentsHandler) GetStudentSubjects(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subject.go"),
			zap.String("function", "GetStudentSubjects"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student model.Student

	if err = json.Unmarshal(reqBytes, &student); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subject.go"),
			zap.String("function", "GetStudentSubjects"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subjects, err := h.subjectService.GetStudentSubjects(context.TODO(), student)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subjects.go"),
			zap.String("function", "GetStudentSubjects()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(subjects)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subjects.go"),
			zap.String("function", "GetStudentSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get student's subject
// @Description get student's subject by id
// @Tags student-subjects
// @Accept json
// @Produce json
// @Param id path int true "ID for getting subject profile"
// @Success 200 {object} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /student/hub/subjects/{id} [get]
func (h *StudentsHandler) GetStudentSubject(w http.ResponseWriter, r *http.Request) {

	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subjects.go"),
			zap.String("function", "GetStudentSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subject, err := h.subjectService.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subjects.go"),
			zap.String("function", "GetStudentSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(subject)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_subjects.go"),
			zap.String("function", "GetStudentSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
