package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Dial to the WebSocket server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to WebSocket server:", err)
		return
	}
	defer conn.Close()

	// Perform WebSocket handshake
	err = performHandshake(conn, "localhost:8080", "/ws")
	if err != nil {
		fmt.Println("Error during handshake:", err)
		return
	}

	fmt.Println("WebSocket connection established.")

	// Start reading messages from the server
	go readMessages(conn)

	// Send a sample message to the server
	sendMessage(conn, "Hello, WebSocket!")

	// Keep the main goroutine running
	select {}
}

func generateWebSocketKey() string {
	// Generate a 16-byte random key and encode it in Base64
	keyBytes := make([]byte, 16)
	_, err := rand.Read(keyBytes)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(keyBytes)
}

func performHandshake(conn net.Conn, host, path string) error {
	key := generateWebSocketKey()
	request := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: Upgrade\r\nUpgrade: websocket\r\nSec-WebSocket-Key: %s\r\nSec-WebSocket-Version: 13\r\n\r\n", path, host, key)
	_, err := conn.Write([]byte(request))
	if err != nil {
		return err
	}

	// Read the response and check if the handshake was successful
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	if !strings.Contains(response, " 101 ") {
		return fmt.Errorf("WebSocket handshake failed. Response: %s", response)
	}

	return nil
}

func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(fmt.Sprintf("Data: %s\r\n", message)))
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString(' ')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Println("Received message:", message)
	}
}
