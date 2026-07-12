package models

// ChatRequest matches the OpenAI chat completions request body shape.
type ChatRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Temperature      *float64      `json:"temperature,omitempty"`
	TopP             *float64      `json:"top_p,omitempty"`
	MaxTokens        *int          `json:"max_tokens,omitempty"`
	Stream           bool          `json:"stream,omitempty"`
	Stop             []string      `json:"stop,omitempty"`
	PresencePenalty  *float64      `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64      `json:"frequency_penalty,omitempty"`
	User             string        `json:"user,omitempty"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

// ChatResponse matches the non-streaming OpenAI chat completions response.
type ChatResponse struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
	Usage   *Usage       `json:"usage,omitempty"`
}

type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason,omitempty"`
}

// TokenChunk matches the streaming chat completion chunk format.
type TokenChunk struct {
	ID      string        `json:"id"`
	Object  string        `json:"object"`
	Created int64         `json:"created"`
	Model   string        `json:"model"`
	Choices []TokenChoice `json:"choices"`
}

type TokenChoice struct {
	Index        int              `json:"index"`
	Delta        ChatMessageDelta `json:"delta"`
	FinishReason *string          `json:"finish_reason"`
}

type ChatMessageDelta struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
