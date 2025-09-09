package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mathaono/freight-simulator/pkg/logger"
	"go.uber.org/zap"
)

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := uuid.New().String()

		// Adiciona ID da request no contexto
		r = r.WithContext(r.Context())

		logger.Info("New Request",
			zap.String("Request ID", reqID),
			zap.String("Method", r.Method),
			zap.String("Path", r.URL.Path),
			zap.String("Remote ADDR", r.RemoteAddr),
			zap.String("User Agent", r.UserAgent()),
		)

		// Executa o próximo handler
		next.ServeHTTP(w, r)

		logger.Info("Response",
			zap.String("Request ID", reqID),
			zap.String("Path", r.URL.Path),
			zap.Duration("latency_ms", time.Since(start)),
		)
	})
}
