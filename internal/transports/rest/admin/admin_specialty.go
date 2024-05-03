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

func (h *AdminsHandler) InitAdminSpecialtiesRoutes(admin *mux.Router) {
	specialties := admin.PathPrefix("/specialties").Subrouter()
	{
		specialties.HandleFunc("", h.CreateSpecialty).Methods(http.MethodPost)
		specialties.HandleFunc("", h.GetSpecialty).Methods(http.MethodGet)
		specialties.HandleFunc("/{id:[0-9]+}", h.GetSpecialty).Methods(http.MethodGet)
		specialties.HandleFunc("/{id:[0-9]+}", h.UpdateSpecialty).Methods(http.MethodPatch)
		specialties.HandleFunc("/{id:[0-9]+}", h.DeleteSpecialty).Methods(http.MethodDelete)
	}
}

// @Summary Create specialty
// @Description create specialty
// @Tags admin-specialties
// @Accept json
// @Produce json
// @Param specialty body model.CreateSpecialtyInput true "Data for creating specialty"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/specialties [post]
func (h *AdminsHandler) CreateSpecialty(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "createSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var specialty model.CreateSpecialtyInput

	if err = json.Unmarshal(reqBytes, &specialty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "createSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Specialties.Create(context.TODO(), specialty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "createSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get specialties
// @Description get specialties
// @Tags admin-specialties
// @Accept json
// @Produce json
// @Success 200 {array} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/specialties [get]
func (h *AdminsHandler) GetSpecialties(w http.ResponseWriter, r *http.Request) {
	specialties, err := h.services.Specialties.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "getSpecialties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(specialties)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "getSpecialties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get specialty
// @Description get specialty by id
// @Tags admin-specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for getting specialty"
// @Success 200 {object} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/specialties/{id} [get]
func (h *AdminsHandler) GetSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	specialty, err := h.services.Specialties.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(specialty)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Update specialty
// @Description update specialty
// @Tags admin-specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for updating specialty"
// @Param request body model.UpdateSpecialtyInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/specialties/{id} [patch]
func (h *AdminsHandler) UpdateSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "updateSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "updateSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var specialty model.UpdateSpecialtyInput
	if err := json.Unmarshal(reqBytes, &specialty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "updateSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.Specialties.Update(context.TODO(), id, specialty)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "updateSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete specialty
// @Description delete specialty
// @Tags admin-specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting specialties"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/specialties/{id} [delete]
func (h *AdminsHandler) DeleteSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "deleteSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Specialties.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "deleteSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
