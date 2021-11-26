package broker

import (
	"encoding/json"
	"fmt"

	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/service"
)

func (k *Kafka) readPayload(payload interface{}, target interface{}) error {
	jsonString, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling event value to json string: %w", err)
	}

	err = json.Unmarshal(jsonString, &target)
	if err != nil {
		return fmt.Errorf("error unmarshaling event value to []byte: %w", err)
	}

	return nil
}

func (k *Kafka) createAccount(payload interface{}, service *service.Service) {
	var account models.Account
	err := k.readPayload(payload, &account)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := service.Authorizer.CreateAccount(models.Account{
		PublicId: account.PublicId,
		Name:     account.Name,
		Username: account.Username,
		Token:    account.Token,
		Role:     account.Role,
	})
	if err != nil {
		fmt.Println("error created accoint in task:", err.Error())
	} else {
		fmt.Println("task.account created", id, account.PublicId, account.Name, account.Username,
			account.Token, account.Role)
	}
}

func (k *Kafka) updateAccount(payload interface{}, service *service.Service) {
	var data models.UpdateAccountInput
	err := k.readPayload(payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = service.Authorizer.UpdateAccount(models.UpdateAccountInput{
		PublicId: data.PublicId,
		Name:     data.Name,
		Password: data.Password,
		Role:     data.Role,
	})
	if err != nil {
		fmt.Println("error update accoint in task:", err.Error())
	} else {
		fmt.Println("task.account updated", data.PublicId, data.Name, data.Password,
			data.Role)
	}
}

func (k *Kafka) deleteAccount(payload interface{}, service *service.Service) {
	var data models.DeleteAccountInput
	err := k.readPayload(payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = service.Authorizer.DeleteAccountByPublicId(data.PublicId)
	if err != nil {
		fmt.Println("error delete accoint in task:", err.Error())
	} else {
		fmt.Println("task.account deleted", data.PublicId)
	}
}
