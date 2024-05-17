package rest

import "github.com/gorilla/mux"

func (h *Handler) initEmployeeRoutes(router *mux.Router) {
	employees := router.PathPrefix("/employee_panel").Subrouter()
	{
		employees.Use(h.authEmployeeMiddleware)
	}
}
