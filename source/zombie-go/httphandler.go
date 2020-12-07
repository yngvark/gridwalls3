package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"log"
	"net/http"
)

type HTTPHandler struct {
	upgrader             *websocket.Upgrader
	network              *network.Broadcaster
	connection           *websocket.Conn
	stopGamelogicChannel chan bool
}

func NewHTTPHandler(allowedOrigins map[string]bool, network *network.Broadcaster, stopGamelogicChannel chan bool) *HTTPHandler {
	h := &HTTPHandler{
		network:              network,
		stopGamelogicChannel: stopGamelogicChannel,
	}

	h.upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin, ok := r.Header["Origin"]
			if !ok {
				return false
			}

			if len(origin) > 0 {
				_, ok := allowedOrigins[origin[0]]
				log.Printf("Checking origin %s. Result: %t\n", origin[0], ok)

				return ok
			}

			return true
		},
		EnableCompression: true,
	}

	return h
}

func (h *HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	connection, err := h.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// This only support one client.
	h.connection = connection

	log.Println("Client connected!")

	// Handle disconnection
	activelyCloseConnectionChannel := make(chan bool)

	defer h.CloseConnectionWhenDone(activelyCloseConnectionChannel)

	go h.ReadIncomingMessages(activelyCloseConnectionChannel)
}

func (h *HTTPHandler) CloseConnectionWhenDone(activelyCloseConnectionChannel chan bool) {
	select {
	case <-h.stopGamelogicChannel:
	case <-activelyCloseConnectionChannel:
	}

	fmt.Println("Closing connection from server")

	err := h.connection.Close()

	if err != nil {
		log.Println("error when closing connection: %w", err)
	} else {
		log.Println("Connection closed successfully.")
	}
}

func (h *HTTPHandler) ReadIncomingMessages(activelyCloseConnectionChannel chan bool) {
	for {
		log.Println("Reading next message...")

		_, message, err := h.connection.ReadMessage()
		if err != nil {
			// Client disconnected
			fmt.Println("Client disconnected")

			// We need to stop both game logic and disconnect
			h.stopGamelogicChannel <- true
			activelyCloseConnectionChannel <- true

			log.Println("Read error:", err)

			break
		}

		h.HandleIncomingMsg(message)
	}
}

func (h *HTTPHandler) HandleIncomingMsg(message []byte) {
	log.Printf("Received: %s", message)
	msgString := string(message)
	h.network.NotifyListeneres(msgString)
}

func (h *HTTPHandler) SendMsg(msg string) error {
	if h.connection == nil {
		return errors.New("could not send message, not connected")
	}

	log.Printf("Sending msg: %s", msg)

	err := h.connection.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("write:", err)
		return err
	}

	return nil
}
