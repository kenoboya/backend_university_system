package rest

import (
	"net/http"
	"test-crud/pkg/logger"

	"go.uber.org/zap"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			zap.String("method", r.Method),
			zap.String("request_url", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}
