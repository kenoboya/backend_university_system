package guest

import (
	"context"
	"encoding/json"
	"net/http"
	"test-crud/internal/transports/rest/common"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *GuestsHandler) InitGuestSpecialtiesRoutes(faculty *mux.Router) {
	faculty.HandleFunc("/{faculty_id}/specialties", h.GetSpecialties).Methods(http.MethodGet)
	faculty.HandleFunc("/{faculty_id}/specialties/{id:[0-9]+}", h.GetSpecialty).Methods(http.MethodGet)
}

// @Summary Get specialties
// @Description get specialties
// @Tags guest-specialties
// @Accept json
// @Produce json
// @Param faculty_id path string true "ID for getting faculty"
// @Success 200 {array} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faculties/{faculty_id}/specialties [get]
func (h *GuestsHandler) GetSpecialties(w http.ResponseWriter, r *http.Request) {
	facultyID, err := common.GetIdStringFromRequest(r, "faculty_id")
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
			zap.String("function", "getSpecialties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	specialties, err := h.specialtiesService.GetSpecialtiesByFacultyID(context.TODO(), facultyID)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
			zap.String("function", "getSpecialties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(specialties)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
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
// @Param faculty_id path string true "ID for getting faculty"
// @Param id path int true "ID for getting specialty"
// @Success 200 {object} model.Specialty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faculties/{faculty_id}/specialties/{id} [get]
func (h *GuestsHandler) GetSpecialty(w http.ResponseWriter, r *http.Request) {
	specialtyID, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	specialty, err := h.specialtiesService.GetById(context.TODO(), uint16(specialtyID))
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(specialty)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_specialty.go"),
			zap.String("function", "getSpecialty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
