package broker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/repository"
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
		//"debug":                "generic,broker,security",
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
		//"debug":                           "generic,broker,security",
	})
	if err != nil {
		return nil, fmt.Errorf("error in kafka constructor, while create consumer: %w", err)
	}

	return &Kafka{
		Producer: producer,
		Consumer: consumer,
	}, nil
}

func (k *Kafka) Subscribe(repos *repository.Repository) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env variables: %s\n", err.Error())
	}

	topic := os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "accounts-stream"
	fmt.Printf("kafka topic %s subscribed", topic)

	err := k.Consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Println("Subscribe kafka ERROR:", err.Error())
	}

	for {
		msg, err := k.Consumer.ReadMessage(-1)
		if err == nil {
			var authAccount models.TaskAccount // что-то общее надо отправлять/ловить, а не модель authe
			err := json.Unmarshal(msg.Value, &authAccount)
			if err != nil {
				fmt.Println("Unmarshal error while decode kafka mess:", err.Error())
			}
			// into gorutine
			accountId, err := repos.CreateAccount(models.TaskAccount{
				PublicId: authAccount.Id,
				Name:     authAccount.Name,
				Username: authAccount.Username,
				Token:    authAccount.Token,
			})
			if err != nil {
				fmt.Println("error created accoint in task:", err.Error())
			}
			fmt.Println(authAccount.Name, authAccount.Username, authAccount.Token)
			fmt.Printf("created accoint in task-service with inner id=%d and public_id=%d\n", accountId, authAccount.Id)

		} else {
			fmt.Printf("Process event from ERROR: %v (%v)\n", err, msg) // TODO логировать
		}
	}
}
