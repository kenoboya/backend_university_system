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
		faculties.HandleFunc("/{faculty_id:[0-9]+}", h.DeleteFaculty).Methods(http.MethodDelete)
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

// @Summary Delete faculty
// @Description delete faculty
// @Tags admin-faculties
// @Accept json
// @Produce json
// @Param faculty_id path int true "ID for deleting faculty"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/faculties/{faculty_id} [delete]
func (h *AdminsHandler) DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdStringFromRequest(r, "faculty_id")
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
