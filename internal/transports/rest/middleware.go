package rest

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: [%s] URL: %s Time: %s", r.Method, r.URL, time.Now().Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}
