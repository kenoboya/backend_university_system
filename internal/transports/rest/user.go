package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) initUsersRoutes(router *mux.Router) {
	router.HandleFunc("/sign-up", h.Users.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", h.Users.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/refresh", h.Users.Refresh).Methods(http.MethodGet)
	user := router.PathPrefix("/users").Subrouter()
	{
		user.Use(h.authMiddleware)
		h.Users.InitUserComplaintsRoutes(user)
		h.Users.InitUserPeopleRoutes(user)
	}
}
