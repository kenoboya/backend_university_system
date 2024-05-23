package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *UsersHandler) InitUserComplaintsRoutes(users *mux.Router) {
	complaints := users.PathPrefix("/complaints").Subrouter()
	{
		complaints.HandleFunc("", h.SubmitComplaint).Methods(http.MethodPost)
	}
}

// @Summary create complaint
// @Description create complaint
// @Tags user-complaints
// @Accept json
// @Produce json
// @Param complaint body model.CreateComplaintInput true "Data for creating complaint"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/complaints [post]
func (h *UsersHandler) SubmitComplaint(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_complaint.go"),
			zap.String("function", "submitComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var complaint model.CreateComplaintInput

	if err = json.Unmarshal(reqBytes, &complaint); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_complaint.go"),
			zap.String("function", "submitComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.complaintsService.Create(context.TODO(), complaint); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_complaint.go"),
			zap.String("function", "submitComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
