package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"zombie-go/pkg/network"
	"zombie-go/pkg/zombie"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := flag.String("addr", "localhost:8080", "http service address")

	fmt.Println("Running")
	flag.Parse()
	log.SetFlags(0)

	broadcaster := network.NewBroadcaster()
	stopGamelogicChannel := make(chan bool)

	httpHandler := NewHTTPHandler(broadcaster, stopGamelogicChannel)
	http.Handle("/zombie", httpHandler)

	var messageSender network.MessageSender = httpHandler
	gameLogic := zombie.NewGameLogic(messageSender, stopGamelogicChannel)

	broadcaster.AddListener(gameLogic)
	// gameLogic.Run(stopGamelogicChannel)

	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}

type HTTPHandler struct {
	upgrader             *websocket.Upgrader
	network              *network.Broadcaster
	connection           *websocket.Conn
	stopGamelogicChannel chan bool
}

func NewHTTPHandler(network *network.Broadcaster, stopGamelogicChannel chan bool) *HTTPHandler {
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

			return len(origin) > 0 && origin[0] == "http://localhost:3000"
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
	//<-done

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

			// We need to stop both game logic and
			activelyCloseConnectionChannel <- true
			h.stopGamelogicChannel <- true

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
