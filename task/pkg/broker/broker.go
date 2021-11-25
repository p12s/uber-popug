package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/service"
)

type Kafka struct {
	Consumer        *kafka.Consumer
	Producer        *kafka.Producer
	TopicAccountBE  string
	TopicAccountCUD string
	TopicTaskBE     string
	TopicTaskCUD    string
	TopicBillingBE  string
	TopicBillingCUD string
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
	} else {
		fmt.Println("task kafka producer 👍")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"metadata.broker.list": os.Getenv("CLOUDKARAFKA_BROKERS"),
		"security.protocol":    "SASL_SSL",
		"sasl.mechanisms":      "SCRAM-SHA-256",
		"sasl.username":        os.Getenv("CLOUDKARAFKA_USERNAME"),
		"sasl.password":        os.Getenv("CLOUDKARAFKA_PASSWORD"),
		"group.id":             os.Getenv("CLOUDKARAFKA_GROUP_ID"),
		"auto.offset.reset":    "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("error in kafka constructor, while create consumer: %w", err)
	} else {
		fmt.Println("task kafka consumer 👍")
	}

	return &Kafka{
		Producer:        producer,
		Consumer:        consumer,
		TopicAccountBE:  os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "account",
		TopicAccountCUD: os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "stream",
		TopicTaskBE:     os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "task",
		TopicTaskCUD:    os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "stream",
		TopicBillingBE:  os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "billing",
		TopicBillingCUD: os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "stream",
	}, nil
}

// TODO сделать хорошее логирование ошибок - stderr (возможно не только ошибок)
// TODO точно ли нужен channel - возможно упростить?
func (k *Kafka) Event(evetType models.EventType, eventTopic string, eventPayload interface{}) {
	deliveryChan := make(chan kafka.Event)

	var data bytes.Buffer
	if err := json.NewEncoder(&data).Encode(models.Event{
		Type:  evetType,
		Value: eventPayload,
	}); err != nil {
		fmt.Printf("task brocker data encode: %s\n", err.Error())
		return
	}

	err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &eventTopic,
			Partition: kafka.PartitionAny,
		},
		Value: data.Bytes(),
	}, deliveryChan)
	if err != nil {
		fmt.Printf("task broker produce: %s\n", err.Error())
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

func (k *Kafka) Subscribe(service *service.Service) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env variables: %s\n", err.Error())
	}

	topics := []string{
		k.TopicAccountBE, k.TopicAccountCUD,
		k.TopicTaskBE, k.TopicBillingBE,
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err := k.Consumer.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Println("Subscribe kafka ERROR:", err.Error())
	} else {
		fmt.Println("task kafka subscribed 👍")
	}

	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := k.Consumer.ReadMessage(1 * time.Second)
			if err != nil {
				continue
			}
			fmt.Printf("✅ Message on %s:\nvalue: %s\n", ev.TopicPartition, string(ev.Value)) // TODO удалить вывод после реализации/обкатки всех событий
			var eventData models.Event
			err = json.Unmarshal(ev.Value, &eventData)
			if err != nil {
				fmt.Println("Unmarshal error:", err.Error())
				return
			}
			k.processEvent(eventData, service)
		}
	}

	fmt.Printf("Closing consumer\n")
	k.Consumer.Close()
}

func (k *Kafka) processEvent(event models.Event, service *service.Service) {
	switch event.Type {
	case models.EVENT_ACCOUNT_CREATED:
		k.createAccount(event.Value, service)
	case models.EVENT_ACCOUNT_UPDATED:
		k.updateAccount(event.Value, service)
	case models.EVENT_ACCOUNT_TOKEN_UPDATED:
		fmt.Println("NEED account token updated event")
	case models.EVENT_ACCOUNT_DELETED:
		k.deleteAccount(event.Value, service)

	default:
		fmt.Println("unknown event type")
	}
}
