package guest

import "test-crud/internal/service"

type GuestsHandler struct {
	facultiesService   service.Faculties
	specialtiesService service.Specialties
}

func NewGuestsHandler(facultiesService service.Faculties, specialtiesService service.Specialties) *GuestsHandler {
	return &GuestsHandler{
		facultiesService:   facultiesService,
		specialtiesService: specialtiesService,
	}
}
