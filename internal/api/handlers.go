package api

import "net/http"


func HandleChatStream(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Write([]byte("data: Here is a piece of data!\n\n"))
	
	if f, ok := w.(http.Flusher); ok {
		f.Flush() // Flush the headers to the client
	}
	

}