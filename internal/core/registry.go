// Package core contains the core application logic, including the model registry.
package core

import (
	"fmt"
	"sse-multiplexer/internal/providers"
)

// ModelRegistry holds a mapping of model names to their corresponding ProviderClient implementations.
type ModelRegistry struct {
	clients map[string]providers.ProviderClient
}

// New creates a new ModelRegistry instance.
func New() *ModelRegistry {
	return &ModelRegistry{
		clients: make(map[string]providers.ProviderClient),
	}
}

// Register adds a new model and its corresponding ProviderClient to the registry.
func (r *ModelRegistry) Register(model string, client providers.ProviderClient) {
	r.clients[model] = client
}

// Get retrieves the ProviderClient for the given model, or returns an error if not found.
func (r *ModelRegistry) Get(model string) (providers.ProviderClient, error) {
	client, ok := r.clients[model]
	if !ok {
		return nil, fmt.Errorf("model %q is not supported or configured", model)
	}
	return client, nil
}
