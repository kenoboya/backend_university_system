package student

import (
	"context"
	"encoding/json"
	"net/http"
	"test-crud/internal/transports/rest/common"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *StudentsHandler) InitStudentProfileRoutes(hub *mux.Router) {
	hub.HandleFunc("/teachers/profile/{id:[0-9]+}", h.GetTeacherProfile).Methods(http.MethodGet)
	hub.HandleFunc("/students/profile/{id:[0-9]+}", h.GetStudentProfile).Methods(http.MethodGet)
}

// @Summary Get teacher profile
// @Description get teacher profile by id
// @Tags student-profiles
// @Accept json
// @Produce json
// @Param id path int true "ID for getting teacher profile"
// @Success 200 {object} model.TeacherBriefInfo "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /student/hub/teachers/profile/{id} [get]
func (h *StudentsHandler) GetTeacherProfile(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetTeacherProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	teacher, err := h.teacherService.GetTeacherProfile(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetTeacherProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(teacher)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetTeacherProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get student profile
// @Description get student profile by id
// @Tags student-profiles
// @Accept json
// @Produce json
// @Param id path int true "ID for getting student profile"
// @Success 200 {object} model.StudentBriefInfo "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /student/hub/students/profile/{id} [get]
func (h *StudentsHandler) GetStudentProfile(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetStudentProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student, err := h.studentService.GetStudentProfile(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetStudentsProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(student)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/student"),
			zap.String("file", "student_profile.go"),
			zap.String("function", "GetStudentsProfile()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
