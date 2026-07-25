package main

import (
	"log/slog"
	"net/http"
	"os"
	"sse-multiplexer/internal/api"
	"sse-multiplexer/internal/core"
	"sse-multiplexer/internal/providers"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	apiKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	if apiKey == "" {
		logger.Error("OPENAI_API_KEY is not set")
		return
	}

	openAIClient := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	openAIAdapter := providers.NewOpenAIAdapter(openAIClient)

	reg := core.New()

	// Register OpenAI models with the registry
	reg.Register("gpt-4o", openAIAdapter)
	reg.Register("gpt-4o-mini", openAIAdapter)

	h := api.NewHandler(logger, reg)

	mux := http.NewServeMux()
	mux.Handle("/chat/completions", api.RequestLogger(logger)(http.HandlerFunc(h.HandleChatStream)))

	port := ":8080"
	logger.Info("Starting Agentgateway SSE Multiplexer on http://localhost"+port, slog.String("port", port))
	logger.Info("Try it out: curl -N -X POST http://localhost"+port+"/chat/completions", slog.String("port", port))

	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Error("Server failed", slog.String("error", err.Error()))
	}
}
