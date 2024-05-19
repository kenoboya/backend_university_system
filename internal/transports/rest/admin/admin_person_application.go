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

func (h *AdminsHandler) InitAdminPeopleRequestsRoutes(application *mux.Router) {
	people := application.PathPrefix("/people").Subrouter()
	{
		people.HandleFunc("", h.GetPeopleApplications).Methods(http.MethodGet)
		people.HandleFunc("/{id:[0-9]+}", h.GetPersonApplication).Methods(http.MethodGet)
		people.HandleFunc("/{id:[0-9]+}", h.ResponseToApplication).Methods(http.MethodPatch)
	}
}

// @Summary Get people applications
// @Description receive applications to create a person
// @Tags admin-applications
// @Accept json
// @Produce json
// @Success 200 {array} model.PersonApplication "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/applications/people [get]
func (h *AdminsHandler) GetPeopleApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := h.services.People.GetListApplications(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "GetPeopleApplications()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(applications)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "GetPeopleApplications()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get person application
// @Description receive applications to create a person
// @Tags admin-applications
// @Accept json
// @Produce json
// @Param id path int true "Person ID for getting application on person"
// @Success 200 {object} model.PersonApplication "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/applications/people/{id} [get]
func (h *AdminsHandler) GetPersonApplication(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "GetPersonApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	application, err := h.services.People.GetApplication(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "GetPersonApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(application)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "GetPersonApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Response to application people
// @Description Response to application people. Accepted or Denied
// @Tags admin-applications
// @Accept json
// @Produce json
// @Param request body model.PeopleApplication true "Response to application"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/applications/people/{id} [patch]
func (h *AdminsHandler) ResponseToApplication(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "ResponseToApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var response model.PersonApplication
	if err := json.Unmarshal(reqBytes, &response); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "ResponseToApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.People.ResponseToApplication(context.TODO(), response); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person_request.go"),
			zap.String("function", "ResponseToApplication()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response.Accepted {
		if err := h.services.Users.ChangeRole(context.TODO(), response.Role, response.UserID); err != nil {
			zap.S().Error(
				zap.String("package", "transport/rest/admin"),
				zap.String("file", "admin_person_request.go"),
				zap.String("function", "ResponseToApplication()"),
				zap.Error(err),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
