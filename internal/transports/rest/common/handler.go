package common

import (
	"context"
	"encoding/json"
	"net/http"
	"test-crud/internal/service"

	"go.uber.org/zap"
)

func GetStudentProfile(w http.ResponseWriter, r *http.Request, studentService service.Students) {
	id, err := GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	student, err := studentService.GetStudentProfile(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(student)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

func GetTeacherProfile(w http.ResponseWriter, r *http.Request, teacherService service.Teachers) {
	id, err := GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	teacher, err := teacherService.GetTeacherProfile(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(teacher)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/common"),
			zap.String("file", "handler.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
