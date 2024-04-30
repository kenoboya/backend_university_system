package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"go.uber.org/zap"
)

// @Summary create faculty
// @Description create faculty
// @Tags faculties
// @Accept json
// @Produce json
// @Param faculty body model.CreateFacultyInput true "Data for creating faculty"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/faculties [post]
func (h *AdminsHandler) createFaculty(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "createFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var faculty model.CreateFacultyInput
	if err := json.Unmarshal(reqBytes, &faculty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "createFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Faculties.Create(context.TODO(), faculty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "createFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get faculties
// @Description get faculties
// @Tags faculties
// @Accept json
// @Produce json
// @Success 200 {array} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/faculties [get]
func (h *AdminsHandler) getFaculties(w http.ResponseWriter, r *http.Request) {
	faculties, err := h.services.Faculties.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "getFaculties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(faculties)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "getFaculties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get faculty
// @Description get faculty by id
// @Tags faculties
// @Accept json
// @Produce json
// @Param id path int true "ID for getting faculty"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/faculties/{id} [get]
func (h *AdminsHandler) getFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	faculty, err := h.services.Faculties.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(faculty)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Delete faculty
// @Description delete faculty
// @Tags faculties
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting faculty"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/faculties/{id} [delete]
func (h *AdminsHandler) deleteFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "deleteFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Faculties.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "deleteFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
