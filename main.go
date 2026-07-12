package main

import (
	"log/slog"
	"net/http"
	"os"
	"sse-multiplexer/internal/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", api.HandleChatStream)

	port := ":8080"
	logger.Info("Starting Agentgateway SSE Multiplexer on http://localhost"+port, slog.String("port", port))
	logger.Info("Try it out: curl -N -X POST http://localhost"+port+"/v1/chat/completions", slog.String("port", port))
	
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Server failed", slog.String("error", err.Error()))
	}
}
