package main

import (
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
) {
	mux.Handle("POST /api/v1/greeting/", handleGreeting(logger))
}
