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

func (h *AdminsHandler) InitAdminFacultiesRoutes(admin *mux.Router) {
	faculties := admin.PathPrefix("/faculties").Subrouter()
	{
		faculties.HandleFunc("", h.CreateFaculty).Methods(http.MethodPost)
		faculties.HandleFunc("", h.GetFaculties).Methods(http.MethodGet)
		faculties.HandleFunc("/{id:[0-9]+}", h.GetFaculty).Methods(http.MethodGet)
		faculties.HandleFunc("/{id:[0-9]+}", h.DeleteFaculty).Methods(http.MethodDelete)
	}
}

// @Summary create faculty
// @Description create faculty
// @Tags admin-faculties
// @Accept json
// @Produce json
// @Param faculty body model.CreateFacultyInput true "Data for creating faculty"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/faculties [post]
func (h *AdminsHandler) CreateFaculty(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
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
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "createFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Faculties.Create(context.TODO(), faculty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
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
// @Tags admin-faculties
// @Accept json
// @Produce json
// @Success 200 {array} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/faculties [get]
func (h *AdminsHandler) GetFaculties(w http.ResponseWriter, r *http.Request) {
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
			zap.String("package", "transport/rest/admin"),
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
// @Tags admin-faculties
// @Accept json
// @Produce json
// @Param id path int true "ID for getting faculty"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/faculties/{id} [get]
func (h *AdminsHandler) GetFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
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
			zap.String("package", "transport/rest/admin"),
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
			zap.String("package", "transport/rest/admin"),
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
// @Tags admin-faculties
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting faculty"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/faculties/{id} [delete]
func (h *AdminsHandler) DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "deleteFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Faculties.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_faculty.go"),
			zap.String("function", "deleteFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
