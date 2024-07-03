package main

import (
	"log/slog"
	"net/http"

	"github.com/jsmithdenverdev/go-rest-microservice/cmd-packages/routes"
)

func newServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	routes.AddRoutes(mux, logger)
	var handler http.Handler = mux
	// handler is now free to be wrapped in additional middleware
	return handler
}
