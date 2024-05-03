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

func (h *Handler) initAdminTeachersRoutes(admin *mux.Router) {
	teachers := admin.PathPrefix("/teachers").Subrouter()
	{
		teachers.HandleFunc("", h.Admins.createTeacher).Methods(http.MethodPost)
		teachers.HandleFunc("", h.Admins.getTeachers).Methods(http.MethodGet)
		teachers.HandleFunc("/{id:[0-9]+}", h.Admins.getTeacher).Methods(http.MethodGet)
		teachers.HandleFunc("/{id:[0-9]+}", h.Admins.updateTeacher).Methods(http.MethodPatch)
		teachers.HandleFunc("/{id:[0-9]+}", h.Admins.deleteTeacher).Methods(http.MethodDelete)
	}
}

// @Summary Create teacher
// @Description create teacher
// @Tags admin-teachers
// @Accept json
// @Produce json
// @Param teacher body model.CreateTeacherInput true "Data for creating teacher"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/teachers [post]
func (h *AdminsHandler) createTeacher(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "createTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var teacher model.CreateTeacherInput

	if err := json.Unmarshal(reqBytes, &teacher); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "createTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Teachers.Create(context.TODO(), teacher); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "createTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get teachers
// @Description get teachers
// @Tags admin-teachers
// @Accept json
// @Produce json
// @Success 200 {array} model.Teacher "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/teachers [get]
func (h *AdminsHandler) getTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.services.Teachers.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "getTeachers()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(teachers)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "getTeachers()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get teacher
// @Description get teacher by id
// @Tags admin-teachers
// @Accept json
// @Produce json
// @Param id path int true "ID for getting teacher"
// @Success 200 {object} model.Teacher "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/teachers/{id} [get]
func (h *AdminsHandler) getTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	teacher, err := h.services.Teachers.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(teacher)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "getTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Update teacher
// @Description update teacher
// @Tags admin-teachers
// @Accept json
// @Produce json
// @Param id path int true "ID for updating teacher"
// @Param request body model.UpdateTeacherInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/teachers/{id} [patch]
func (h *AdminsHandler) updateTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "updateTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "updateTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var teacher model.UpdateTeacherInput
	if err := json.Unmarshal(reqBytes, &teacher); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "updateTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.Teachers.Update(context.TODO(), id, teacher); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "updateTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete teacher
// @Description delete teacher
// @Tags admin-teachers
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting teacher"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/teachers/{id} [delete]
func (h *AdminsHandler) deleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "deleteTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.Teachers.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_teacher.go"),
			zap.String("function", "deleteTeacher()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
