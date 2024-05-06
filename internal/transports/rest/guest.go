package rest

import (
	"github.com/gorilla/mux"
)

func (h *Handler) initGuestsRoutes(guest *mux.Router) {
	faculties := h.Guests.InitGuestFacultiesRoutes(guest)
	h.Guests.InitGuestSpecialtiesRoutes(faculties)
}
