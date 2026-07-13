package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sse-multiplexer/internal/models"
)

func HandleChatStream(w http.ResponseWriter, r *http.Request) {

	
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	log.Println("Handler received request")
	// Parse request body to ChatRequest type
	var requestBody models.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	log.Println("Request body ", requestBody)
	w.Write([]byte("data: Here is a piece of data!\n\n"))

	// if f, ok := w.(http.Flusher); ok {
	// 	f.Flush() // Flush the headers to the client
	// }

	http.ResponseWriter.WriteHeader(w, http.StatusOK)

}
