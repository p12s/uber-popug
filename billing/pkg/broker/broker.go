package broker

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/p12s/uber-popug/billing/pkg/models"
	"github.com/p12s/uber-popug/billing/pkg/repository"
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

func (k *Kafka) Subscribe(repos *repository.Repository) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env variables: %s\n", err.Error())
	}

	kafkaPrefix := os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX")
	// сделал названия топиков по комменту: https://github.com/p12s/uber-popug/pull/2#discussion_r748330569
	topics := []string{
		kafkaPrefix + "accounts",
		kafkaPrefix + "task-lifecycle",
		kafkaPrefix + "billing-transactions",
		kafkaPrefix + "stream",
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err := k.Consumer.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Println("Subscribe kafka ERROR:", err.Error())
	} else {
		fmt.Println("Subscribed kafka OK")
	}

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-k.Consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				k.Consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				k.Consumer.Unassign()
			case *kafka.Message:
				//fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				var eventData models.Event
				err := json.Unmarshal(e.Value, &eventData)
				if err != nil {
					fmt.Println("error:", err.Error())
					// TODO сохранить в лог ошибку
					return
				}
				processEvent(eventData, repos)

			case kafka.PartitionEOF:
				//fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				// fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	k.Consumer.Close()
}

func processEvent(eventData models.Event, repos *repository.Repository) {
	switch eventData.Type {
	case models.EVENT_ACCOUNT_CREATED:
		var account models.Account

		// TODO тут баг с анмаршалингом - непраильно достаются поля, поправить
		fmt.Println("before unmarshal")
		fmt.Println(eventData.Value)

		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		err := enc.Encode(eventData.Value)
		if err != nil {
			fmt.Println("ERR", err.Error())
		}

		err = json.Unmarshal(buffer.Bytes(), &account)
		if err != nil {
			fmt.Println("Unmarshal error while decode kafka mess:", err.Error())
		}
		createAccount(account, repos)
	case models.EVENT_TASK_CREATED:
		fmt.Println("task created event")
	case models.TASK:
		fmt.Println("task assigned event")
	case models.EVENT_TASK_COMPLETED:
		fmt.Println("task completed event")
	}
}

func createAccount(account models.Account, repos *repository.Repository) {
	fmt.Println("catched marshaled account:")
	fmt.Println(account)
	// TODO улушить сохранение ошибки
	id, err := repos.CreateAccount(models.Account{
		PublicId:  account.PublicId,
		Name:      account.Name,
		Username:  account.Username,
		Token:     account.Token,
		Role:      account.Role,
		CreatedAt: account.CreatedAt,
	})
	if err != nil {
		fmt.Println("error created account in task:", err.Error())
	} else {
		fmt.Println("task.account created", id, account.PublicId, account.Name, account.Username,
			account.Token, account.Role)
	}
}

func createTask(task models.Task, repos *repository.Repository) {
	fmt.Println("catched marshaled task:")
	fmt.Println(task)
	// TODO улушить сохранение ошибки
	id, err := repos.CreateTask(models.Task{
		PublicId:          task.PublicId,
		AssignedAccountId: task.AssignedAccountId,
		Description:       task.Description,
		JiraId:            task.JiraId,
		Status:            task.Status,
		CreatedAt:         task.CreatedAt,
	})
	if err != nil {
		fmt.Println("error created task in billing:", err.Error())
	} else {
		fmt.Println("billing.task created", id, task.PublicId, task.AssignedAccountId, task.Description,
			task.JiraId, task.Status, task.CreatedAt)
	}
}
