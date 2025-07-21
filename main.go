package main

import (
	"chat-app-backend/handlers"
	"chat-app-backend/models"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var Client = make(map[*websocket.Conn]*models.User)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
