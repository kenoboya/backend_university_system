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

func (h *Handler) initAdminGroupsRoutes(admin *mux.Router) {
	groups := admin.PathPrefix("/groups").Subrouter()
	{
		groups.HandleFunc("", h.Admins.createGroup).Methods(http.MethodPost)
		groups.HandleFunc("", h.Admins.getGroups).Methods(http.MethodGet)
		groups.HandleFunc("/{id:[0-9]+}", h.Admins.getGroups).Methods(http.MethodGet)
		groups.HandleFunc("/{id:[0-9]+}", h.Admins.deleteGroup).Methods(http.MethodDelete)
	}
}

// @Summary create group
// @Description create group
// @Tags admin-groups
// @Accept json
// @Produce json
// @Param faculty body model.CreateGroupInput true "Data for creating group"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/groups [post]
func (h *AdminsHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "createGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var group model.CreateGroupInput
	if err := json.Unmarshal(reqBytes, &group); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "createGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Groups.Create(context.TODO(), group); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "createGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get groups
// @Description get groups
// @Tags admin-groups
// @Accept json
// @Produce json
// @Success 200 {array} model.Group "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/groups [get]
func (h *AdminsHandler) getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.services.Groups.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "getGroups()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(groups)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "getGroups()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get group
// @Description get group by id
// @Tags admin-groups
// @Accept json
// @Produce json
// @Param id path int true "ID for getting group"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/groups/{id} [get]
func (h *AdminsHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "getGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	group, err := h.services.Groups.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "getGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(group)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "getGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Delete group
// @Description delete group
// @Tags admin-groups
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting group"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/groups/{id} [delete]
func (h *AdminsHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "deleteGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Groups.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_group.go"),
			zap.String("function", "deleteGroup()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
