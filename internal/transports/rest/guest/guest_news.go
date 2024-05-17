package guest

import (
	"context"
	"encoding/json"
	"net/http"
	"test-crud/internal/transports/rest/common"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *GuestsHandler) InitGuestNewsRoutes(guest *mux.Router) *mux.Router {
	news := guest.PathPrefix("/news").Subrouter()
	{
		news.HandleFunc("", h.GetListNews).Methods(http.MethodGet)
		news.HandleFunc("/{id:[0-9]+}", h.GetNews).Methods(http.MethodGet)
	}
	return news
}

// @Summary get news
// @Description get news
// @Tags guest-news
// @Accept json
// @Produce json
// @Success 200 {array} model.News "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news [get]
func (h *GuestsHandler) GetListNews(w http.ResponseWriter, r *http.Request) {
	news, err := h.newsService.GetList(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_news.go"),
			zap.String("function", "GetListNews()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(news)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_news.go"),
			zap.String("function", "GetListNews()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary get new
// @Description get new by id
// @Tags guest-news
// @Accept json
// @Produce json
// @Param id path string true "ID for getting news"
// @Success 200 {object} model.Faculty "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news/{id} [get]
func (h *GuestsHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_news.go"),
			zap.String("function", "GetNews()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	news, err := h.newsService.GetNews(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_news.go"),
			zap.String("function", "GetNews()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(news)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest/guest"),
			zap.String("file", "guest_news.go"),
			zap.String("function", "GetNews()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}
