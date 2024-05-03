package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) initUsersRoutes(users *mux.Router) {
	users.HandleFunc("/sign-up", h.Users.signUp).Methods(http.MethodPost)
	users.HandleFunc("/sign-in", h.Users.signIn).Methods(http.MethodPost)
	users.HandleFunc("/refresh", h.Users.refresh).Methods(http.MethodGet)
}
