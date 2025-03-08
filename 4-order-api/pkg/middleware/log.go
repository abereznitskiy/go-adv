package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"status":   wrapper.StatusCode,
			"method":   r.Method,
			"path":     r.URL.Path,
			"duration": time.Since(start).Milliseconds(),
		}).Info("Request processed")
	})
}
