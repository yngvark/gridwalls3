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

func main() {
	fmt.Println("Running")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/zombie", zombie)
	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}
