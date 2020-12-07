package main

import (
	"flag"
	"fmt"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/mainhelp"
	"log"
	"net/http"
	"os"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/zombie"
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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalf("env variable PORT is not set")
	}

	printAllowedCorsOrigins(allowedCorsOrigins)

	setupGame(allowedCorsOrigins)

	listenAndServe(port)
}

func printAllowedCorsOrigins(allowedCorsOrigins map[string]bool) {
	fmt.Println("ALLOWED_CORS_ORIGINS:")
	for k := range allowedCorsOrigins {
		fmt.Printf("- %s\n", k)
	}
	fmt.Println()
}

func setupGame(allowedCorsOrigins map[string]bool) {
	broadcaster := network.NewBroadcaster()
	stopGamelogicChannel := make(chan bool)

	httpHandler := NewHTTPHandler(allowedCorsOrigins, broadcaster, stopGamelogicChannel)
	http.Handle("/zombie", httpHandler)

	var messageSender network.MessageSender = httpHandler
	gameLogic := zombie.NewGameLogic(messageSender, stopGamelogicChannel)

	broadcaster.AddListener(gameLogic)
}

func listenAndServe(port string) {
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		out := []byte("OK")
		_, err := writer.Write(out)

		fmt.Println("error when responding on /health: %s", err)
	})

	serverAddr := flag.String("addr", fmt.Sprintf(":%s", port), "http service address")
	log2().Infof("Running on %s\n", *serverAddr)
	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}
