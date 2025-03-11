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
		userAgent := r.Header.Get("User-Agent")
		next.ServeHTTP(wrapper, r)
		duration := time.Since(start)
		log.WithFields(log.Fields{
			"method":       r.Method,
			"url":          r.URL.Path,
			"code":         wrapper.StatusCode,
			"path":         r.URL.Path,
			"statusCode":   wrapper.StatusCode,
			"duration":     duration.String(),
			"responseSize": wrapper.ResponseSize,
			"userAgent":    userAgent,
			"requestID":    r.Context().Value("requestID"),
		}).Info("Request processed")
	})
}
