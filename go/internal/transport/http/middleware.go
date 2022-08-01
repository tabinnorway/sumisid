package http

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"][0]
		if r != nil && r.Header != nil && r.Header["Authorization"] != nil {
			println("=== AUTH: ", authHeader)
		}
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		} else {
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		}
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("Starting request")
		next.ServeHTTP(w, r)
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("Finished request")
	})
}
