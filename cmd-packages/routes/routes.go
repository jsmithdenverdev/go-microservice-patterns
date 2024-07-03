package routes

import (
	"log/slog"
	"net/http"

	"github.com/jsmithdenverdev/go-rest-microservice/cmd-packages/handlers"
)

func AddRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
) {
	mux.Handle("POST /api/v1/greeting/", handlers.HandleGreeting(logger))
}
