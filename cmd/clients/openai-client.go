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

// SimpleChat sends one non-streaming chat completion request and returns text.
func SimpleChat(ctx context.Context, prompt string) (string, error) {
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

	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: defaultModel,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("empty choices in OpenAI response")
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)
	if content == "" {
		return "", fmt.Errorf("empty message content in OpenAI response")
	}

	return content, nil
}

func main() {
	prompt := flag.String("prompt", "Say hello in one short sentence.", "Prompt to send to /v1/chat/completions")
	timeout := flag.Duration("timeout", 30*time.Second, "Request timeout")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	out, err := SimpleChat(ctx, *prompt)
	if err != nil {
		log.Fatalf("chat request failed: %v", err)
	}

	fmt.Println(out)
}
