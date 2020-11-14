package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"zombie-go/pkg/network"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := flag.String("addr", "localhost:8080", "http service address")

	fmt.Println("Running")
	flag.Parse()
	log.SetFlags(0)

	h := NewHTTPHandler()
	http.Handle("/zombie", h)

	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}

type HTTPHandler struct {
	Upgrader *websocket.Upgrader
	Network  *network.Network
}

func NewHTTPHandler() *HTTPHandler {
	h := &HTTPHandler{}

	h.Upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin, ok := r.Header["Origin"]
			if !ok {
				return false
			}

			return len(origin) > 0 && origin[0] == "http://localhost:3000"
		},
		EnableCompression: true,
	}

	h.Network = network.New()

	rec := TestReceiver{}
	h.Network.AddListener(rec)

	return h
}

type TestReceiver struct {
}

func (t TestReceiver) MsgReceived(msg string) {
	panic("implement me")
}

//func (t *TestReceiver) MsgReceived(msg string) {
//	fmt.Printf("TestReceiver recived: %s\n", msg)
//}

func (h *HTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	connection, err := h.Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Println("Client connected!")

	// Handle disconnection
	connectionCloserChannel := make(chan bool)
	defer h.CloseConnectionWhenDone(connection, connectionCloserChannel)

	go h.ReadIncomingMessages(connection, connectionCloserChannel)
}

func (h *HTTPHandler) CloseConnectionWhenDone(connection *websocket.Conn, done chan bool) {
	<-done
	err := connection.Close()
	if err != nil {
		log.Println("error when closing connection: %w", err)
	} else {
		log.Println("Connection closed successfully.")
	}
}

func (h *HTTPHandler) ReadIncomingMessages(connection *websocket.Conn, connectionCloserChannel chan bool) {
	for {
		log.Println("Reading next message...")
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			// Client disconnected
			connectionCloserChannel <- true
			log.Println("Read error:", err)
			break
		}

		h.HandleIncomingMsg(message, messageType, connection)
	}
}

func (h *HTTPHandler) HandleIncomingMsg(message []byte, messageType int, connection *websocket.Conn) {
	log.Printf("Received: %s", message)
	msgString := string(message)
	h.Network.MsgReceived(msgString)
}
