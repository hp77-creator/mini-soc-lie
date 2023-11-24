package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

func main() {
	serverURL := "ws://localhost:8080/ws"
	message := "hello"

	// Create a WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Client connected")

	// Send the message to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	// Receive and print the server's response
	_, response, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Error reading server response:", err)
		os.Exit(1)
	}

	fmt.Printf("Server Response: %s\n", response)

	// Gracefully close the connection
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		fmt.Println("Error sending close message:", err)
		os.Exit(1)
	}

	// Wait for a short time before exiting
	time.Sleep(1 * time.Second)
}
