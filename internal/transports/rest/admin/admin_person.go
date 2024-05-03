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

func (h *AdminsHandler) InitAdminPeopleRoutes(admin *mux.Router) {
	people := admin.PathPrefix("/people").Subrouter()
	{
		people.HandleFunc("", h.CreatePerson).Methods(http.MethodPost)
		people.HandleFunc("", h.GetPeople).Methods(http.MethodGet)
		people.HandleFunc("/{id:[0-9]+}", h.GetPerson).Methods(http.MethodGet)
		people.HandleFunc("/{id:[0-9]+}", h.UpdatePerson).Methods(http.MethodPatch)
		people.HandleFunc("/{id:[0-9]+}", h.DeletePerson).Methods(http.MethodDelete)
	}
}

// @Summary create person
// @Description create person
// @Tags admin-people
// @Accept json
// @Produce json
// @Param person body model.CreatePersonInput true "Data for creating person"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/people [post]
func (h *AdminsHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "createPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var person model.CreatePersonInput
	if err := json.Unmarshal(reqBytes, &person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "createPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.People.Create(context.TODO(), person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "createPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get people
// @Description get people
// @Tags admin-people
// @Accept json
// @Produce json
// @Success 200 {array} model.Person "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/people [get]
func (h *AdminsHandler) GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := h.services.People.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "getPeople()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(people)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "getPeople()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get person
// @Description get person by id
// @Tags admin-people
// @Accept json
// @Produce json
// @Param id path int true "ID for getting person"
// @Success 200 {object} model.Person "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/people/{id} [get]
func (h *AdminsHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "user_person.go"),
			zap.String("function", "getPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := h.services.People.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "getPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "getPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Update person
// @Description update person
// @Tags admin-people
// @Accept json
// @Produce json
// @Param id path int true "ID for updating person"
// @Param request body model.UpdatePersonInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/people/{id} [patch]
func (h *AdminsHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "user_person.go"),
			zap.String("function", "updatePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "user_person.go"),
			zap.String("function", "updatePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var person model.UpdatePersonInput
	if err := json.Unmarshal(reqBytes, &person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "updatePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.People.Update(context.TODO(), id, person)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "updatePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete person
// @Description delete person
// @Tags admin-people
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting person"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/people/{id} [delete]
func (h *AdminsHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "deletePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.People.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/admin"),
			zap.String("file", "admin_person.go"),
			zap.String("function", "deletePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
