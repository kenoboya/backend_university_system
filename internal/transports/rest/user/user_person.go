package user

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
// @Tags user-person
// @Accept json
// @Produce json
// @Param person body model.CreatePersonInput true "Data for creating complaint"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/people [post]
func (h *UsersHandler) SubmitPerson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_person.go"),
			zap.String("function", "SubmitPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var person model.CreatePersonInput
	if err = json.Unmarshal(reqBytes, &person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_person.go"),
			zap.String("function", "SubmitPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.peopleService.CreateApplicationPerson(context.TODO(), person); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/user"),
			zap.String("file", "user_person.go"),
			zap.String("function", "SubmitPerson()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
