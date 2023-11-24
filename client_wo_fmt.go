package main

// without fmt

import (
	"github.com/gorilla/websocket"
	"os"
)

func fwrite(text []byte) {
	f, _ := os.OpenFile("/dev/stdout", os.O_WRONLY, 0)
	f.Write(text)
	f.Close()
}

func main() {
	serverURL := "ws://localhost:8080/ws"
	message := "hello"

	// Create a WebSocket connection
	conn, _, _ := websocket.DefaultDialer.Dial(serverURL, nil)
	//if err != nil {
	//	string text := "Error connecting to WebSocket:" + err.Error()
	//	fwrite(text)
	//	os.Exit(1)
	//}
	defer conn.Close()

	//fmt.Println("Client connected")

	// Send the message to the server
	_ = conn.WriteMessage(websocket.TextMessage, []byte(message))
	//if err != nil {
	//	fmt.Println("Error sending message:", err)
	//	os.Exit(1)
	//}

	// Receive and print the server's response
	_, response, _ := conn.ReadMessage()
	//if err != nil {
	//	fmt.Println("Error reading server response:", err)
	//	os.Exit(1)
	//}

	//fmt.Printf("Server Response: %s\n", response)
	fwrite(response)

	// Gracefully close the connection
	_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//if err != nil {
	//	fmt.Println("Error sending close message:", err)
	//	os.Exit(1)
	//}

	// Wait for a short time before exiting
	//time.Sleep(1 * time.Second)
}
