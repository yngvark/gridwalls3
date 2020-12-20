package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	gamelogicPkg "github.com/yngvark/gridwalls3/source/zombie-go/pkg/gamelogic"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/log2"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/pubsub"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/pulsar_connector"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(fmt.Errorf("could not get logger: %w\n", err))
	}

	// CONSUME ---------------------------------

	//consume()

	//allowedCorsOrigins, err := m.GetAllowedCorsOrigins(os.LookupEnv, "ALLOWED_CORS_ORIGINS")
	//if err != nil {
	//	logger.Fatalf("could not get cors env: %s", err)
	//}
	//
	//port, ok := os.LookupEnv("PORT")
	//if !ok {
	//	logger.Fatalf("env variable PORT is not set")
	//}

	//m.PrintAllowedCorsOrigins(allowedCorsOrigins)
	//m.SetupGame(allowedCorsOrigins)
	//m.HttpListen(port, logger)

	//fmt.Println("--- WAITING ---------------------------------------------------------")
	//<- doneSignal
}

func run() error {
	logger, err := log2.New()
	if err != nil {
	    return fmt.Errorf("could not create logger: %w", err)
	}

	broker := pubsub.NewBroker()
	stopGamelogicChannel := make(chan bool)

	pulsarProducer, err := pulsar_connector.NewPulsarPublisher(logger, context.Background())
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	defer pulsarProducer.Close()

	var publisher pubsub.Publisher = pulsarProducer
	gameLogic := gamelogicPkg.NewGameLogic(logger, publisher, stopGamelogicChannel)

	var subscriber pubsub.Subscriber = gameLogic
	broker.Subscribe(subscriber)

	gameLogic.Run()
	return nil
}

func consume() {
	fmt.Println("--- CONSUME ---------------------------------------------------------")
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:36650",
	})

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "zombie",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})

	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
		msg.ID(), string(msg.Payload()))
}
