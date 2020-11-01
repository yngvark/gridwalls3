package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func generate(connection *websocket.Conn, done chan bool) {
	fmt.Println("Starting to generate...")
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Println("Zombie generator stopped.")
			return
		case <-ticker.C:
			log.Println("Sending message...")
			err := connection.WriteMessage(websocket.TextMessage, []byte("Here is a string...."))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
