package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const defaultModel = openai.ChatModelGPT4oMini
const localOpenAIBaseURL = "http://localhost:8080"

// SimpleChatStream sends a streaming chat completion request and prints chunks.
func SimpleChatStream(ctx context.Context, prompt string) (string, error) {
	apiKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	if strings.TrimSpace(prompt) == "" {
		prompt = "Say hello in one short sentence."
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(localOpenAIBaseURL),
	)

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Model: defaultModel,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})

	var contentBuilder strings.Builder
	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) == 0 {
			continue
		}

		delta := chunk.Choices[0].Delta.Content
		if delta == "" {
			continue
		}

		fmt.Print(delta)
		contentBuilder.WriteString(delta)
	}

	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("create chat completion stream: %w", err)
	}

	content := strings.TrimSpace(contentBuilder.String())
	if content == "" {
		return "", fmt.Errorf("empty message content in OpenAI stream response")
	}

	return content, nil
}

func main() {
	prompt := flag.String("prompt", "Say hello in one short sentence.", "Prompt to send to /v1/chat/completions")
	timeout := flag.Duration("timeout", 30*time.Second, "Request timeout")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	_, err := SimpleChatStream(ctx, *prompt)
	if err != nil {
		log.Fatalf("chat request failed: %v", err)
	}

	fmt.Println()
}
