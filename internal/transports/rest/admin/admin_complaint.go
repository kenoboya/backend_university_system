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

func (h *AdminsHandler) InitAdminComplaintsRoutes(admin *mux.Router) {
	complaints := admin.PathPrefix("/complaints").Subrouter()
	{
		complaints.HandleFunc("", h.GetComplaints).Methods(http.MethodGet)
		complaints.HandleFunc("/{id:[0-9]+}", h.GetComplaint).Methods(http.MethodGet)
		complaints.HandleFunc("/{id:[0-9]+}", h.ResponseToComplaint).Methods(http.MethodPatch)
	}
}

// @Summary Get Complaints
// @Description get complaints
// @Tags admin-complaints
// @Accept json
// @Produce json
// @Success 200 {array} model.Complaint "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/complaints [get]
func (h *AdminsHandler) GetComplaints(w http.ResponseWriter, r *http.Request) {
	complaints, err := h.services.Complaints.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "getComplaints()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(complaints)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "getComplaints()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get complaint
// @Description get complaint by id
// @Tags admin-complaints
// @Accept json
// @Produce json
// @Param id path int true "ID for getting complaint"
// @Success 200 {object} model.Complaint "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/complaints/{id} [get]
func (h *AdminsHandler) GetComplaint(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "getComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	complaint, err := h.services.Complaints.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "getComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(complaint)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "getComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Response complaint
// @Description Response to complaint
// @Tags admin-complaints
// @Accept json
// @Produce json
// @Param id path int true "ID for response complaint"
// @Param request body model.ResponseComplaintInput true "Response to complaint"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/complaints/{id} [patch]
func (h *AdminsHandler) ResponseToComplaint(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "ResponseToComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "ResponseToComplaint"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var response model.ResponseComplaintInput
	if err := json.Unmarshal(reqBytes, &response); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "ResponseToComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.Complaints.Response(context.TODO(), id, response); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_complaint.go"),
			zap.String("function", "ResponseToComplaint()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
