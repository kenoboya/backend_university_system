package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"go.uber.org/zap"
)

// @Summary create person
// @Description create person
// @Tags people
// @Accept json
// @Produce json
// @Param person body model.CreatePersonInput true "Data for creating person"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/people [post]
func (h *UsersHandler) createPerson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "createPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var person model.CreatePersonInput
	if err := json.Unmarshal(reqBytes, &person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "createPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.peopleService.Create(context.TODO(), person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
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
// @Tags people
// @Accept json
// @Produce json
// @Success 200 {array} model.Person "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/people [get]
func (h *UsersHandler) getPeople(w http.ResponseWriter, r *http.Request) {
	people, err := h.peopleService.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "getPeople()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(people)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
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
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "ID for getting person"
// @Success 200 {object} model.Person "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/people/{id} [get]
func (h *UsersHandler) getPerson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "getPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := h.peopleService.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "getPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
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
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "ID for updating person"
// @Param request body model.UpdatePersonInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/people/{id} [patch]
func (h *UsersHandler) updatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
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
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "updatePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.peopleService.Update(context.TODO(), id, person)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
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
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting person"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /user/people/{id} [delete]
func (h *UsersHandler) deletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "deletePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.peopleService.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "user_person.go"),
			zap.String("function", "deletePerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
