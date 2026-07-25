package providers

import (
	"context"
	"sse-multiplexer/internal/models"
)

// providers/interface.go
type ProviderClient interface {
    Stream(ctx context.Context, req *models.ChatRequest) (<-chan string, <-chan error)
}