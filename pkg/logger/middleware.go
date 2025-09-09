package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Middleware(log *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Infow("New Request",
				"Method", r.Method,
				"Path", r.URL.Path,
				"Remote", r.RemoteAddr,
			)

			next.ServeHTTP(w, r)

			log.Infow("Response",
				"Path", r.URL.Path,
				"latency_ms", time.Since(start).Seconds()*1000,
			)
		})
	}
}
