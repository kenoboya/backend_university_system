package teacher

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

func (h *TeachersHandler) InitTeacherLessonsRoutes(hub *mux.Router) {
	lessons := hub.PathPrefix("/lessons").Subrouter()
	{
		lessons.HandleFunc("/{id:[0-9]+}", h.GetLesson).Methods(http.MethodGet)
		lessons.HandleFunc("/{id:[0-9]+}/attendance", h.AttendanceList).Methods(http.MethodGet)
		lessons.HandleFunc("/{id:[0-9]+}/attendance", h.MarkAttendance).Methods(http.MethodPatch)
		lessons.HandleFunc("/{id:[0-9]+}/grade", h.GradeList).Methods(http.MethodGet)
		lessons.HandleFunc("/{id:[0-9]+}/grade", h.Evaluate).Methods(http.MethodPatch)
	}
}

// @Summary Get lesson
// @Description get lesson by id
// @Tags teacher-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID for getting lesson"
// @Success 200 {object} model.Lesson "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/lessons/{id} [get]
func (h *TeachersHandler) GetLesson(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lesson, err := h.lessonService.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(lesson)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lesson.go"),
			zap.String("function", "getLesson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Attendance List of students
// @Description Attendance List of students
// @Tags teacher-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID for getting lesson"
// @Success 200 {array} model.StudentBriefInfo "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/lessons/{id}/attendance [get]
func (h *TeachersHandler) AttendanceList(w http.ResponseWriter, r *http.Request) {
	lesson_id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "AttendanceList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	attendanceRecords, err := h.studentService.GetStudentsAttendance(context.TODO(), lesson_id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "AttendanceList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(attendanceRecords)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "AttendanceList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Mark student's attendance
// @Description Mark student's attendance
// @Tags teacher-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID lesson"
// @Param request body model.AttendanceRecord true "information for mark attendance"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/lessons/{id}/attendance [patch]
func (h *TeachersHandler) MarkAttendance(w http.ResponseWriter, r *http.Request) {
	lesson_id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "MarkAttendance()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "MarkAttendance()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var attendanceRecord model.AttendanceRecord
	if err := json.Unmarshal(reqBytes, &attendanceRecord); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "MarkAttendance()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.teacherService.MarkAttendance(context.TODO(), lesson_id, attendanceRecord)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "MarkAttendance()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Grade List of students
// @Description Grade List of students
// @Tags teacher-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID for getting lesson"
// @Success 200 {array} model.StudentBriefInfo "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/lessons/{id}/grade [get]
func (h *TeachersHandler) GradeList(w http.ResponseWriter, r *http.Request) {
	lesson_id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "GradeList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	grades, err := h.studentService.GetStudentsGrades(context.TODO(), lesson_id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "GradeList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(grades)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "GradeList()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary evaluate the student
// @Description evaluate the student
// @Tags teacher-lessons
// @Accept json
// @Produce json
// @Param id path int true "ID lesson"
// @Param request body model.StudentBriefInfo true "Student information for evaluate"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /teacher/hub/lessons/{id}/grade [patch]
func (h *TeachersHandler) Evaluate(w http.ResponseWriter, r *http.Request) {
	lesson_id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "Evaluate()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "Evaluate()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var grade model.Grade
	if err := json.Unmarshal(reqBytes, &grade); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "Evaluate()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.teacherService.EvaluteStudent(context.TODO(), lesson_id, grade)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/teacher"),
			zap.String("file", "teacher_lessons.go"),
			zap.String("function", "Evaluate()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
