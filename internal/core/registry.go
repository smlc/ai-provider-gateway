// core/registry.go
package core

import (
	"fmt"
	"sse-multiplexer/internal/providers"
)

type ModelRegistry struct {
    clients map[string]providers.ProviderClient
}

func New() *ModelRegistry {
    return &ModelRegistry{
        clients: make(map[string]providers.ProviderClient),
    }
}

func (r *ModelRegistry) Register(model string, client providers.ProviderClient) {
    r.clients[model] = client
}

func (r *ModelRegistry) Get(model string) (providers.ProviderClient, error) {
    client, ok := r.clients[model]
    if !ok {
        return nil, fmt.Errorf("model %q is not supported or configured", model)
    }
    return client, nil
}