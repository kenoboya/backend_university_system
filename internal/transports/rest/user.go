package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"
	"test-crud/internal/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UsersHandler struct {
	service service.Users
}

func NewUsersHandler(service service.Users) *UsersHandler {
	return &UsersHandler{service: service}
}

func (h *UsersHandler) initRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		users.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)
	}
}
func (h *UsersHandler) signUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Fatal(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user model.UserSignUpInput
	if err := json.Unmarshal(reqBytes, &user); err != nil {
		zap.S().Fatal(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Validator(); err != nil {
		zap.S().Fatal(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.service.SignUp(context.TODO(), user); err != nil {
		zap.S().Fatal(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *UsersHandler) signIn(w http.ResponseWriter, r *http.Request) {
	// todo
}
