package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"go.uber.org/zap"
)

// @Summary create group
// @Description create group
// @Tags groups
// @Accept json
// @Produce json
// @Param faculty body model.CreateGroupInput true "Data for creating group"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/groups [post]
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
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} model.Group "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/groups [get]
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
// @Tags groups
// @Accept json
// @Produce json
// @Param id path int true "ID for getting group"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/groups/{id} [get]
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
// @Tags groups
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting group"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/groups/{id} [delete]
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
