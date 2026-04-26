// internal/handler/handler.go
package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var reqTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"path", "status"},
)

func Health(w http.ResponseWriter, r *http.Request) {
	reqTotal.WithLabelValues("/health", "200").Inc()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		slog.Error("encode error", "err", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	slog.Info("hello request", "remote", r.RemoteAddr)
	reqTotal.WithLabelValues("/hello", "200").Inc()
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"}); err != nil {
		slog.Error("encode error", "err", err)
	}
}
