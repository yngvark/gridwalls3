package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/mainhelp"
	"log"
	"net/http"
	"os"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/zombie"

	"github.com/gorilla/websocket"
)

func main() {
	err := InitLogger()
	if err != nil {
	    log.Fatalf("could not get logger: %w", err)
	}

	allowedCorsOrigins, err := mainhelp.GetAllowedCorsOrigins(os.LookupEnv, "ALLOWED_CORS_ORIGINS")
	if err != nil {
		log.Fatalf("could not get cors env: %s", err)
	}

	fmt.Println("ALLOWED_CORS_ORIGINS:")
	for k := range allowedCorsOrigins {
		fmt.Printf("- %s\n", k)
	}
	fmt.Println()
	flag.Parse()
	log.SetFlags(0)

	broadcaster := network.NewBroadcaster()
	stopGamelogicChannel := make(chan bool)

	httpHandler := NewHTTPHandler(allowedCorsOrigins, broadcaster, stopGamelogicChannel)
	http.Handle("/zombie", httpHandler)

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		out := []byte("OK")
		_, err := writer.Write(out)

		fmt.Println("error when responding on /health: %s", err)
	})

	var messageSender network.MessageSender = httpHandler
	gameLogic := zombie.NewGameLogic(messageSender, stopGamelogicChannel)

	broadcaster.AddListener(gameLogic)
	// gameLogic.Run(stopGamelogicChannel)

	err = listenAndServe()
	if err != nil {
		log.Fatal(": %w", err)
	}
}

func listenAndServe() error {
	certFile, useCert := os.LookupEnv("CERTIFICATE_FILE")
	keyFile, useKey := os.LookupEnv("KEY_FILE")

	if useCert || useKey {
		if !(useCert && useKey) {
			return fmt.Errorf("both CERTIFICATE_FILE and KEY_FILE need to be set")
		}

		port := "8443"
		serverAddr := flag.String("addr", fmt.Sprintf("localhost:%s", port), "http service address")

		log2().Info("Using TLS")
		log2().Infof("Running on %s\n", *serverAddr)
		//log.Println("Using TLS")
		//fmt.Printf("Running on %s\n", *serverAddr)

		log.Fatal(http.ListenAndServeTLS(*serverAddr, certFile, keyFile, nil))
	} else {
		port := "8081"
		serverAddr := flag.String("addr", fmt.Sprintf("localhost:%s", port), "http service address")

		log2().Infof("Running on %s\n", *serverAddr)
		//fmt.Printf("Running on %s\n", *serverAddr)

		log.Fatal(http.ListenAndServe(*serverAddr, nil))
	}

	return nil
}

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
