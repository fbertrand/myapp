// cmd/main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"myapp/internal/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/hello", handler.Hello)
	mux.Handle("/metrics", promhttp.Handler()) // endpoint Prometheus

	slog.Info("server starting", "port", 8080)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
