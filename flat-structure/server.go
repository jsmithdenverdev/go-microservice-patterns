package main

import (
	"log/slog"
	"net/http"
)

func newServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger)
	var handler http.Handler = mux
	// handler is now free to be wrapped in additional middleware
	return handler
}
