package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var mu sync.Mutex

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("WebSocket Server Started on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("WebSocket Server Failed to Start.")
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		mu.Lock()
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, msg)
		}
		mu.Unlock()
	}
}
