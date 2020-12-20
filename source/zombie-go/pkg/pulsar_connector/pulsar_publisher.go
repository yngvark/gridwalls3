package pulsar_connector

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"go.uber.org/zap"
	"time"
)

type PulsarPublisher struct {
	log *zap.SugaredLogger
	ctx context.Context
	client pulsar.Client
	producer pulsar.Producer
}

func NewPulsarPublisher(logger *zap.SugaredLogger, ctx context.Context) (*PulsarPublisher, error) {
	p := &PulsarPublisher{
		log: logger,
		ctx: ctx,
	}

	err := p.init()
	return p, err
}

func (m *PulsarPublisher) init() error {
	// Create client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:36650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
	    return fmt.Errorf("could not instantiate Pulsar client: %w", err)
	}

	m.client = client

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "zombie",
	})

	if err != nil {
	   return fmt.Errorf("could not create producer: %w", err)
	}

	m.producer = producer

	return nil
}

func (m *PulsarPublisher) SendMsg(msg string) error {
	_, err := m.producer.Send(m.ctx, &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})

	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (m *PulsarPublisher) Close() {
	m.producer.Close()
	m.client.Close()
}
