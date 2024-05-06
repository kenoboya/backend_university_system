package admin

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

func (h *AdminsHandler) InitAdminStudentsRoutes(admin *mux.Router) {
	students := admin.PathPrefix("/students").Subrouter()
	{
		students.HandleFunc("", h.CreateStudent).Methods(http.MethodPost)
		students.HandleFunc("", h.GetStudents).Methods(http.MethodGet)
		students.HandleFunc("/{id:[0-9]+}", h.GetStudent).Methods(http.MethodGet)
		students.HandleFunc("/{id:[0-9]+}", h.UpdateStudent).Methods(http.MethodPatch)
		students.HandleFunc("/{id:[0-9]+}", h.DeleteStudent).Methods(http.MethodDelete)
	}
}

// @Summary create student
// @Description create student
// @Tags admin-students
// @Accept json
// @Produce json
// @Param student body model.CreateStudentInput true "Data for creating student"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/students [post]
func (h *AdminsHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "createStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var student model.CreateStudentInput

	if err = json.Unmarshal(reqBytes, &student); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_students.go"),
			zap.String("function", "createStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Students.Create(context.TODO(), student); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "createStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get students
// @Description get students
// @Tags admin-students
// @Accept json
// @Produce json
// @Success 200 {array} model.Student "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/students [get]
func (h *AdminsHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.services.Students.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "getStudents()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(students)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "getStudents()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get student
// @Description get student by id
// @Tags admin-students
// @Accept json
// @Produce json
// @Param id path int true "ID for getting student"
// @Success 200 {object} model.Student "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/students/{id} [get]
func (h *AdminsHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	common.GetStudent(w, r, h.services.Students)
}

// @Summary Update student
// @Description update student
// @Tags admin-students
// @Accept json
// @Produce json
// @Param id path int true "ID for updating student"
// @Param request body model.UpdateStudentInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/students/{id} [patch]
func (h *AdminsHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "updateStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "updateStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var student model.UpdateStudentInput
	if err := json.Unmarshal(reqBytes, &student); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "updateStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.Students.Update(context.TODO(), id, student)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "updateStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete student
// @Description delete student
// @Tags admin-students
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting student"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/students/{id} [delete]
func (h *AdminsHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "deleteStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Students.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_student.go"),
			zap.String("function", "deleteStudent()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
