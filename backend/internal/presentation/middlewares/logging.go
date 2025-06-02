package middlewares

import (
	"net/http"
	"time"

	"github.com/repyg/DockerMonitoringApp/backend/pkg/utils"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger utils.LoggerInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			clientIP := utils.GetClientIP(r)

			wrapper := &responseWriterWrapper{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(wrapper, r)
			duration := time.Since(start)

			if wrapper.statusCode >= 200 && wrapper.statusCode < 400 {
				logger.Infof(
					"REQUESTS: %s - %s - %s - %d - %s",
					clientIP,
					r.Method,
					r.URL.Path,
					wrapper.statusCode,
					duration,
				)
			} else {
				logger.Errorf("REQUESTS: %s - %s - %s - %d - %s", clientIP, r.Method, r.URL.Path, wrapper.statusCode, duration)
			}
		})
	}
}
