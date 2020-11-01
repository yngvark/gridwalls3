package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func closeConnectionWhenDone(connection *websocket.Conn, done chan bool) {
	<-done
	err := connection.Close()
	if err != nil {
		log.Println("error when closing connection: %w", err)
	} else {
		log.Println("Connection closed successfully.")
	}
}

func readIncoming(connection *websocket.Conn, listeners []chan bool) {
	for {
		log.Println("Reading next message...")
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			notifyListeners(listeners, true)
			log.Println("Read error:", err)
			break
		}

		err = handleIncomingMsg(message, messageType, connection)
		if err != nil {
			notifyListeners(listeners, true)
			log.Printf("errror when handling incoming message: %w", err)
			return
		}
	}
}

func notifyListeners(listeners []chan bool, signal bool) {
	log.Printf("Notifying listeners. Count: %d\n", len(listeners))
	for _, l := range listeners {
		l <- signal
	}
}
