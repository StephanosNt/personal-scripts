package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow frontend access
	response := Response{Message: "Hello from the Go backend!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/message", handler)
	fmt.Println("Backend running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
