package main

import (
	"fmt"
	"log"
	"net/http"
)

func zombie(writer http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Println("Client connected!")

	listeners := make([]chan bool, 0)

	connectionCloserChannel := make(chan bool)
	listeners = append(listeners, connectionCloserChannel)

	defer closeConnectionWhenDone(connection, connectionCloserChannel)

	generateChannel := make(chan bool)
	listeners = append(listeners, generateChannel)
	fmt.Printf("Len listeners: %d\n", len(listeners))
	go generate(connection, generateChannel)

	go readIncoming(connection, listeners)
}
