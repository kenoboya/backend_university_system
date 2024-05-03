package admin

import (
	"test-crud/internal/service"
)

type AdminsHandler struct {
	services service.Services
}

func NewAdminsHandler(services service.Services) *AdminsHandler {
	return &AdminsHandler{services: services}
}
