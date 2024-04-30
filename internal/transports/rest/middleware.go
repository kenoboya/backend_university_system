package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"test-crud/pkg/logger"

	"go.uber.org/zap"
)

type CtxValue int

const (
	ctxUserID           CtxValue = iota
	authorizationHeader          = "Authorization"
)

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			logger.Error(
				zap.String("method", r.Method),
				zap.String("request_url", r.RequestURI),
				zap.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user_id, err := h.tokenManager.VerifyToken(token)
		if err != nil {
			logger.Error(
				zap.String("method", r.Method),
				zap.String("request_url", r.RequestURI),
				zap.String("error", err.Error()),
			)
		}
		ctx := context.WithValue(r.Context(), ctxUserID, user_id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			zap.String("method", r.Method),
			zap.String("request_url", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		return "", errors.New("auth header is empty")
	}

	headerParts := strings.Split(header, "")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}
	return headerParts[1], nil
}
