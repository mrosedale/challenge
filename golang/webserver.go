package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response structure
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Test    string `json:"test"`
}

// Handler function
func handler(w http.ResponseWriter, r *http.Request) {
	// Create response
	response := Response{
		Message: "Hello, World!",
		Status:  200,
		Test:    "ok",
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode response as JSON and send
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func main() {
	// Set route and handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/status", status)

	// Start server on port 8080
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
