package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"go.uber.org/zap"
)

// @Summary create subject
// @Description create subject
// @Tags subjects
// @Accept json
// @Produce json
// @Param subject body model.CreateSubjectInput true "Data for creating subject"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/subjects [post]
func (h *AdminsHandler) createSubject(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "createSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var subject model.CreateSubjectInput

	if err = json.Unmarshal(reqBytes, &subject); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "createSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Subjects.Create(context.TODO(), subject); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "createSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get subjects
// @Description get subjects
// @Tags subjects
// @Accept json
// @Produce json
// @Success 200 {array} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/subjects [get]
func (h *AdminsHandler) getSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := h.services.Subjects.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "getSubjects()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(subjects)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "getSubjects()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get subject
// @Description get subject by id
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path int true "ID for getting subject"
// @Success 200 {object} model.Subject "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/subjects/{id} [get]
func (h *AdminsHandler) getSubject(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "getSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	subject, err := h.services.Subjects.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "getSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(subject)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "getSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Update subject
// @Description update subject
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path int true "ID for updating subject"
// @Param request body model.UpdateSubjectInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/subjects/{id} [patch]
func (h *AdminsHandler) updateSubject(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "updateSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "updateSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var subject model.UpdateSubjectInput
	if err := json.Unmarshal(reqBytes, &subject); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "updateSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.Subjects.Update(context.TODO(), id, subject)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "updateSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete subject
// @Description delete subject
// @Tags subjects
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting subject"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/subjects/{id} [delete]
func (h *AdminsHandler) deleteSubject(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "deleteSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Subjects.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_subject.go"),
			zap.String("function", "deleteSubject()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
