package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"sse-multiplexer/internal/core"
	"sse-multiplexer/internal/logging"
	"sse-multiplexer/internal/models"
)

// Handler holds shared dependencies for the API handlers.
type Handler struct {
	logger   *slog.Logger
	registry *core.ModelRegistry
}

// NewHandler creates a Handler with the provided base logger.
func NewHandler(logger *slog.Logger, registry *core.ModelRegistry) *Handler {
	return &Handler{logger: logger, registry: registry}
}

// HandleChatStream handles POST /chat/completions and streams an SSE response.
func (h *Handler) HandleChatStream(w http.ResponseWriter, r *http.Request) {
	logger := logging.FromContextOr(r.Context(), h.logger)

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	logger.Info("Handler received request")

	// Parse request body to ChatRequest type
	var requestBody models.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error("Failed to decode request body", slog.String("error", err.Error()))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	logger.Info("Request body parsed", slog.String("model", requestBody.Model))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("data: Here is a piece of data!\n\n"))

	// if f, ok := w.(http.Flusher); ok {
	// 	f.Flush() // Flush the headers to the client
	// }
}
