package main

import (
	"golang.org/x/net/websocket"
	"os"
)

func fwrite_c(text string) {
	f, _ := os.OpenFile("/dev/stdout", os.O_WRONLY, 0)
	byteSlice := []byte(text)
	f.Write(byteSlice)
	f.Close()
}

func main() {
	serverUrl := "ws://localhost:8080/ws"
	ws, err := websocket.Dial(serverUrl, "", "http://localhost")
	if err != nil {
		text := "Error connecting to WebSocket:" + err.Error()
		fwrite_c(text)
		os.Exit(1)
	}
	defer ws.Close()

	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			//time.Sleep(1 * time.Second)
			//putting software delay to only send one message to server
			for i := 0; i < 1e7; i++ {
			}
			message := "Hello"
			err := websocket.Message.Send(ws, message)
			if err != nil {
				text := "Error connecting to WebSocket:" + err.Error()
				fwrite_c(text)
				return
			}
			fwrite_c("Sent message: " + message + "\n")
		}
	}()

	for {
		select {
		//case <-interrupt:
		//	fwrite_c("Interrupt signal received. Closing websocket")
		//	return
		default:
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				fwrite_c("Error receiving message: " + err.Error())
				return
			}
			fwrite_c("\nReceived message from server: " + msg + "\n")
			return
		}
	}
}
