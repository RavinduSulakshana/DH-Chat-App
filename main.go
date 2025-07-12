package main

import (
	"chatapp/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServer(":8080", nil)
}
