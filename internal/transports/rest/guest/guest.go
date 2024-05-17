package guest

import "test-crud/internal/service"

type GuestsHandler struct {
	facultiesService   service.Faculties
	specialtiesService service.Specialties
	newsService        service.News
}

func NewGuestsHandler(facultiesService service.Faculties, specialtiesService service.Specialties, newsService service.News) *GuestsHandler {
	return &GuestsHandler{
		facultiesService:   facultiesService,
		specialtiesService: specialtiesService,
		newsService:        newsService,
	}
}
