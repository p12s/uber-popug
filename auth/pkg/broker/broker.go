package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/p12s/uber-popug/auth/pkg/models"
)

type Kafka struct {
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
		fmt.Println("auth kafka producer 👍")
	}

	return &Kafka{
		Producer:        producer,
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
		fmt.Printf("auth brocker data encode: %s\n", err.Error())
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
		fmt.Printf("auth broker produce: %s\n", err.Error())
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
