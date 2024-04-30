package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"go.uber.org/zap"
)

// @Summary Create specialty
// @Description create specialty
// @Tags specialties
// @Accept json
// @Produce json
// @Param specialty body model.CreateSpecialtyInput true "Data for creating specialty"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/specialties [post]
func (h *AdminsHandler) createSpecialty(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "createSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Specialties.Create(context.TODO(), specialty); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
// @Tags specialties
// @Accept json
// @Produce json
// @Success 200 {array} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/specialties [get]
func (h *AdminsHandler) getSpecialties(w http.ResponseWriter, r *http.Request) {
	specialties, err := h.services.Specialties.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
// @Tags specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for getting specialty"
// @Success 200 {object} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/specialties/{id} [get]
func (h *AdminsHandler) getSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
// @Tags specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for updating specialty"
// @Param request body model.UpdateSpecialtyInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/specialties/{id} [patch]
func (h *AdminsHandler) updateSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
// @Tags specialties
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting specialties"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/specialties/{id} [delete]
func (h *AdminsHandler) deleteSpecialty(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "deleteSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Specialties.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_specialty.go"),
			zap.String("function", "deleteSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
