package guest

import (
	"context"
	"encoding/json"
	"net/http"
	"test-crud/internal/transports/rest/common"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *GuestsHandler) InitGuestFacultiesRoutes(guest *mux.Router) *mux.Router {
	faculties := guest.PathPrefix("/faculties").Subrouter()
	{
		faculties.HandleFunc("", h.GetFaculties).Methods(http.MethodGet)
		faculties.HandleFunc("/{id:[0-9]+}", h.GetFaculty).Methods(http.MethodGet)
	}
	return faculties
}

// @Summary Get faculties
// @Description get faculties
// @Tags guest-faculties
// @Accept json
// @Produce json
// @Success 200 {array} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faculties [get]
func (h *GuestsHandler) GetFaculties(w http.ResponseWriter, r *http.Request) {
	faculties, err := h.facultiesService.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_faculty.go"),
			zap.String("function", "getFaculties()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(faculties)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_faculty.go"),
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
// @Tags guest-faculties
// @Accept json
// @Produce json
// @Param id path string true "ID for getting faculty"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /faculties/{id} [get]
func (h *GuestsHandler) GetFaculty(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdStringFromRequest(r, "faculty_id")
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	faculty, err := h.facultiesService.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(faculty)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_faculty.go"),
			zap.String("function", "getFaculty()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
