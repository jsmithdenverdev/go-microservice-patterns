package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func handleGreeting(logger *slog.Logger) http.Handler {
	type GreetingRequest struct {
		Name string `json:"name"`
	}

	type GreetingResponse struct {
		Greeting string `json:"greeting"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := decode[GreetingRequest](r)
		if err != nil {
			logger.ErrorContext(r.Context(), "failed to decode greeting request", "error", err)
			if err := encode(
				w, r,
				http.StatusInternalServerError,
				UnhandledError{Message: ErrorMessageUnhandled}); err != nil {
				logger.ErrorContext(r.Context(), "failed to encode error response", "error", err)
				return
			}
		}

		if req.Name == "" {
			req.Name = "World"
		}

		greeting := fmt.Sprintf("Hello, %s", req.Name)

		if err := encode(w, r,
			http.StatusOK,
			GreetingResponse{Greeting: greeting}); err != nil {
			logger.ErrorContext(r.Context(), "failed to encode response", "error", err)
			return
		}
	})
}
