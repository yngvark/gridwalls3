package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var serverAddr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin, ok := r.Header["Origin"]
		if !ok {
			return false
		}

		return len(origin) > 0 && origin[0] == "http://localhost:3000"
	},
	EnableCompression: true,
}

func zombie(writer http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Println("Client connected!")

	done := make(chan bool)
	defer closeConnectionWhenDone(done, connection)

	go readIncoming(connection, done)
	go generate(connection)
}

func closeConnectionWhenDone(done chan bool, connection *websocket.Conn) {
	<-done
	err := connection.Close()
	if err != nil {
		log.Println("error when closing connection: %w", err)
	}
}

func readIncoming(connection *websocket.Conn, done chan bool) {
	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			done <- true
			log.Println("Read error:", err)
			break
		}

		err = handleIncomingMsg(message, messageType, connection)
		if err != nil {
			done <- true
			log.Printf("errror when handling incoming message: %w", err)
			return
		}
	}
}

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

func generate(connection *websocket.Conn) {
	fmt.Println("Starting to generate...")
	//connection.WriteMessage(websocket.TextMessage, []byte("Here is a string...."))
}

func main() {
	fmt.Println("Running")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/zombie", zombie)
	//http.HandleFunc("/zombie", zombieOld)
	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}

//
//func zombieOld(writer http.ResponseWriter, request *http.Request) {
//	connection, err := upgrader.Upgrade(writer, request, nil)
//	if err != nil {
//		log.Print("upgrade:", err)
//		return
//	}
//
//	defer closeConnection(connection)
//
//	for {
//		messageType, message, err := connection.ReadMessage()
//		if err != nil {
//			log.Println("Read error:", err)
//			break
//		}
//
//		err = handleIncomingMsg(message, messageType, connection)
//		if err != nil {
//			log.Printf("errror when handling incoming message: %w", err)
//			return
//		}
//	}
//
//	fmt.Println("Zombie done")
//}
