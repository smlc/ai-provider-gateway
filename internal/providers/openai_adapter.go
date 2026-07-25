package providers

import (
	"context"
	"fmt"
	"strings"

	"sse-multiplexer/internal/models"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

// OpenAIAdapter adapts the OpenAI SDK client to the ProviderClient interface.
type OpenAIAdapter struct {
	client openai.Client
}

func NewOpenAIAdapter(client openai.Client) *OpenAIAdapter {
	return &OpenAIAdapter{client: client}
}

func (a *OpenAIAdapter) Stream(ctx context.Context, req *models.ChatRequest) (<-chan string, <-chan error) {
	tokens := make(chan string)
	errs := make(chan error, 1)

	go func() {
		defer close(tokens)
		defer close(errs)

		messages, err := toOpenAIMessages(req.Messages)
		if err != nil {
			errs <- err
			return
		}

		if len(messages) == 0 {
			errs <- fmt.Errorf("at least one message is required")
			return
		}

		params := openai.ChatCompletionNewParams{
			Model:    shared.ChatModel(req.Model),
			Messages: messages,
		}

		if req.Temperature != nil {
			params.Temperature = openai.Float(*req.Temperature)
		}
		if req.TopP != nil {
			params.TopP = openai.Float(*req.TopP)
		}
		if req.MaxTokens != nil {
			params.MaxTokens = openai.Int(int64(*req.MaxTokens))
		}
		if req.PresencePenalty != nil {
			params.PresencePenalty = openai.Float(*req.PresencePenalty)
		}
		if req.FrequencyPenalty != nil {
			params.FrequencyPenalty = openai.Float(*req.FrequencyPenalty)
		}
		if strings.TrimSpace(req.User) != "" {
			params.User = openai.String(req.User)
		}

		stream := a.client.Chat.Completions.NewStreaming(ctx, params)
		for stream.Next() {
			chunk := stream.Current()
			for _, choice := range chunk.Choices {
				if choice.Delta.Content == "" {
					continue
				}

				select {
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				case tokens <- choice.Delta.Content:
				}
			}
		}

		if err := stream.Err(); err != nil {
			errs <- fmt.Errorf("openai stream failed: %w", err)
		}
	}()

	return tokens, errs
}

func toOpenAIMessages(messages []models.ChatMessage) ([]openai.ChatCompletionMessageParamUnion, error) {
	out := make([]openai.ChatCompletionMessageParamUnion, 0, len(messages))

	for _, msg := range messages {
		role := strings.ToLower(strings.TrimSpace(msg.Role))
		content := msg.Content

		switch role {
		case "system":
			out = append(out, openai.SystemMessage(content))
		case "developer":
			out = append(out, openai.DeveloperMessage(content))
		case "assistant":
			out = append(out, openai.AssistantMessage(content))
		case "user", "":
			out = append(out, openai.UserMessage(content))
		default:
			return nil, fmt.Errorf("unsupported message role %q", msg.Role)
		}
	}

	return out, nil
}
