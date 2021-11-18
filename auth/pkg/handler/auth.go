package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/auth/pkg/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.Account

	input.PublicId = uuid.New()
	input.Role = models.Employee
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateAccount(input) // достать весь акк, не только ид
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO вынести отсюда и сделать логирование ошибок (возможно не только ошибок)
	// TODO точно ли нужен channel - возможно упростить
	go func() {
		deliveryChan := make(chan kafka.Event)

		var data bytes.Buffer
		if err := json.NewEncoder(&data).Encode(models.Event{
			Type:  models.EVENT_ACCOUNT_CREATED,
			Value: input,
		}); err != nil {
			fmt.Printf("auth brocker data encode: %s\n", err.Error())
			return
		}

		cudStreamTopic := os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "stream"
		err = h.broker.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &cudStreamTopic,
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
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO здесь можно отправить токен в другие сервисы,
	// когда юзер запросит страницу из отдельного сервиса - токен в его браузере == токену в сервисе
	// и не надо будет запрашивать сервис Auth

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) token(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	account, err := h.services.GetAccountById(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}
