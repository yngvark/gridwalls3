package main

import (
	"fmt"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/log2"
	"log"
	"os"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/mainhelp"
)

func main() {
	logger, err := log2.New()
	if err != nil {
		log.Fatal(fmt.Errorf("could not get logger: %w\n", err))
	}

	m := mainhelp.New(logger)

	allowedCorsOrigins, err := m.GetAllowedCorsOrigins(os.LookupEnv, "ALLOWED_CORS_ORIGINS")
	if err != nil {
		logger.Fatalf("could not get cors env: %s", err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		logger.Fatalf("env variable PORT is not set")
	}

	m.PrintAllowedCorsOrigins(allowedCorsOrigins)
	m.SetupGame(allowedCorsOrigins)
	m.HttpListen(port, logger)
}
