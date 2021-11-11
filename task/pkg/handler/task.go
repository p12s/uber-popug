package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/p12s/uber-popug/task/pkg/models"
)

func (h *Handler) getTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	_ = accountId // TODO возможно будут проверки, принадлежит ли этот таск текущему юзеру

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	task, err := h.services.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) createTask(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	_ = accountId // TODO возможно будет использоваться

	var input models.Task
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// пусть пока таск ассайнится на текущего пользователя
	id, err := h.services.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	// TODO вынести отсюда и сделать логирование ошибок (возможно не только ошибок)
	// TODO точно ли нужен channel - возможно упростить
	go func() {
		deliveryChan := make(chan kafka.Event)

		var data bytes.Buffer
		if err := json.NewEncoder(&data).Encode(models.Task{
			Id:                id,
			AssignedAccountId: input.AssignedAccountId,
			Description:       input.Description,
			Status:            input.Status,
		}); err != nil {
			fmt.Printf("auth brocker data encode: %s\n", err.Error())
			return
		}

		accountsStreamTopic := os.Getenv("CLOUDKARAFKA_TOPIC_PREFIX") + "tasks"
		err = h.broker.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &accountsStreamTopic,
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
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
