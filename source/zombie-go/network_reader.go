package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func handleIncomingMsg(message []byte, messageType int, connection *websocket.Conn) error {
	log.Printf("Received: %s", message)
	msgString := string(message)
	reply := "You said: " + msgString

	err := connection.WriteMessage(messageType, []byte(reply))
	if err != nil {
		log.Println("write:", err)
		return err
	}

	return nil
}
