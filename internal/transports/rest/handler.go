package rest

import (
	"net/http"
	_ "test-crud/docs"
	"test-crud/internal/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // URL для Swagger JSON
	))
	students := router.PathPrefix("/students").Subrouter()
	{
		students.HandleFunc("", h.createStudent).Methods(http.MethodPost)
		students.HandleFunc("/{id:[0-9]+}", h.deleteStudent).Methods(http.MethodDelete)
		students.HandleFunc("/{id:[0-9]+}", h.updateStudent).Methods(http.MethodPatch)
		students.HandleFunc("", h.getAllStudents).Methods(http.MethodGet)
		students.HandleFunc("/{id:[0-9]+}", h.getStudentById).Methods(http.MethodGet)
	}
	return router
}
