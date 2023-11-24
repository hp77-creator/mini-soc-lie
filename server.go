package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Customize this function to allow connections from the desired origin(s).
		// For example, you can compare the origin header with a list of allowed origins.
		// In this example, any origin is allowed.
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebsockets(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error connecting socket: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println(messageType)
		fmt.Println(err)
		if err != nil {
			break
		}

		switch messageType {
		case websocket.TextMessage:
			// Handle text message
			fmt.Printf("Received Text Message: %s\n", p)

			// Echo the message back to the client
			err := conn.WriteMessage(messageType, p)
			if err != nil {
				fmt.Println("Error writing message:", err)
				return
			}

		case -1:
			// Handle close message
			fmt.Println("Client initiated WebSocket close")
			return
		}
	}

}

func main() {
	http.HandleFunc("/ws", handleWebsockets)

	fmt.Println("WebSocket server is running at ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
