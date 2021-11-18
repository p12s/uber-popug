package broker

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Consumer *kafka.Consumer
	Producer *kafka.Producer
}

func NewKafka() (*Kafka, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"metadata.broker.list": os.Getenv("CLOUDKARAFKA_BROKERS"),
		"security.protocol":    "SASL_SSL",
		"sasl.mechanisms":      "SCRAM-SHA-256",
		"sasl.username":        os.Getenv("CLOUDKARAFKA_USERNAME"),
		"sasl.password":        os.Getenv("CLOUDKARAFKA_PASSWORD"),
	})
	if err != nil {
		return nil, fmt.Errorf("error in kafka constructor, while create producer: %w", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"metadata.broker.list":            os.Getenv("CLOUDKARAFKA_BROKERS"),
		"security.protocol":               "SASL_SSL",
		"sasl.mechanisms":                 "SCRAM-SHA-256",
		"sasl.username":                   os.Getenv("CLOUDKARAFKA_USERNAME"),
		"sasl.password":                   os.Getenv("CLOUDKARAFKA_PASSWORD"),
		"group.id":                        "cloudkarafka-example",
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})
	if err != nil {
		return nil, fmt.Errorf("error in kafka constructor, while create consumer: %w", err)
	}

	return &Kafka{
		Producer: producer,
		Consumer: consumer,
	}, nil
}
