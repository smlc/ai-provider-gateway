package providers

import (
	"context"
	"sse-multiplexer/internal/models"
)

// ProviderClient defines the interface for streaming providers.
type ProviderClient interface {
	Stream(ctx context.Context, req *models.ChatRequest) (<-chan string, <-chan error)
}
