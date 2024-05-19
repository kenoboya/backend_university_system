package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) initUsersRoutes(users *mux.Router) {
	users.HandleFunc("/sign-up", h.Users.SignUp).Methods(http.MethodPost)
	users.HandleFunc("/sign-in", h.Users.SignIn).Methods(http.MethodPost)
	users.HandleFunc("/refresh", h.Users.Refresh).Methods(http.MethodGet)

	complaints := users.PathPrefix("/complaints").Subrouter()
	{
		complaints.Use(h.authMiddleware)
		complaints.HandleFunc("", h.Users.SubmitComplaint).Methods(http.MethodPost)
	}

	people := users.PathPrefix("/people").Subrouter()
	{
		people.Use(h.authMiddleware)
		people.HandleFunc("", h.Users.SubmitPerson).Methods(http.MethodPost)
	}
}
