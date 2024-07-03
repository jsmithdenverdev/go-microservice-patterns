package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func HandleGreeting(logger *slog.Logger) http.Handler {
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
			http.Error(w, errUnhandled.Error(), http.StatusInternalServerError)
			return
		}

		if req.Name == "" {
			req.Name = "World"
		}

		greeting := fmt.Sprintf("Hello, %s!", req.Name)

		if err := encode(w, r,
			http.StatusOK,
			GreetingResponse{Greeting: greeting}); err != nil {
			logger.ErrorContext(r.Context(), "failed to encode greeting response", "error", err)
			http.Error(w, errUnhandled.Error(), http.StatusInternalServerError)
			return
		}
	})
}
