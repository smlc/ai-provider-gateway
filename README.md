# AI Provider Gateway (`ai-provider-gateway`)

A high-performance, lightweight AI proxy gateway written in Go. This project standardizes streaming AI/LLM completions across multiple model providers (e.g., OpenAI, self-hosted vLLM) behind a unified, OpenAI-compatible streaming interface.

---

## 🏗️ Architecture & Design Principles

Architecture:
* The **HTTP routing layer** knows nothing about how a specific LLM generates text.
* The **LLM client providers** know nothing about HTTP Server-Sent Events (SSE) flushing.

### The 3 Core Layers

1. **Shared Language (`internal/models`)**
   * Establishes a universal format (OpenAI Chat Completions JSON standard).
   * Defines structured types for incoming requests (`ChatRequest`), outgoing streamed responses (`ChatResponse`), and token chunks (`TokenChunk`).

2. **Provider Clients (`internal/providers`)**
   * Implements actual network clients (e.g., OpenAI API, local vLLM, Groq, or mock clients).
   * Adheres to a unified `ProviderClient` interface using Go channels and context-aware streaming.

3. **The Multiplexer (`internal/core`)**
   * Manages the SSE loop, buffer flushing (`http.Flusher`), and client context cancellation.
   * Consumes token channels returned by the `ProviderClient` without coupling to specific provider implementations.

---

## 📁 Repository Structure

```plaintext
/
├── cmd/
│   └── clients/
│       └── openai-clients.go           # Test client
├── internal/
│   ├── api/
│   │   ├── handlers.go       # HTTP handlers (e.g., HandleChatCompletion)
│   │   ├── middleware.go     # Auth checks, rate limiting, and request logging
│   ├── core/
│   │   └── multiplexer.go    # SSE loop, buffer flushing, and context cancellation logic
│   ├── models/
│   │   └── openai.go         # Standardized request & response structs (ChatRequest, TokenChunk)
│   └── providers/
│       ├── interface.go      # The LLMClient contract interface
│       └── openai_adapter.go # Client for OpenAI public API
├── go.mod
└── go.sum
```

---

## 🚀 Key Features

* **Context-Aware Streaming:** Automatically halts model generation and upstream network calls if the client drops the connection (`ctx.Done()`).
* **Zero-Coupling Multiplexing:** Decouples SSE handling from model provider logic.
* **Extensible Provider Architecture:** Easily add support for Anthropic, Groq, Ollama, or custom GPU endpoints by implementing `LLMClient`.

---

## 🗺️ Implementation Roadmap

- [x] **Core Architecture & Directory Layout**
- [ ] **Shared Models Definition:** Create standard structs for OpenAI Chat Completions JSON formats (`ChatRequest`, `TokenChunk`).
- [ ] **Provider Client Implementation:** Build mock/real HTTP streaming clients (vLLM, OpenAI).
- [ ] **HTTP Layer Refactoring:** Move routing and SSE handler logic into `internal/api` and `internal/core`.

---

## 🛠️ Getting Started

### Prerequisites

* **Go**: `1.21+`

### Running the Server

```bash
# Clone the repository
git clone https://github.com/smlc/ai-provider-gateway.git
cd ai-provider-gateway

# Build and run
export OPENAI_API_KEY=test-key 
go run main.go

# Run the test client
export OPENAI_API_KEY=test-key     
go run ./cmd/clients -prompt "hello from standalone client"
```

---