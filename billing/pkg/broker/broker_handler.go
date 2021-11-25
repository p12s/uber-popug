package broker

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/p12s/uber-popug/billing/pkg/models"
	"github.com/p12s/uber-popug/billing/pkg/service"
)

const (
	MIN_PRICE = -10
	ZERO      = 0
	MAX_PRICE = 20
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
		fmt.Println("error created accoint in billing:", err.Error())
	} else {
		fmt.Println("billing.account created", id, account.PublicId, account.Name, account.Username,
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
		fmt.Println("error update accoint in billing:", err.Error())
	} else {
		fmt.Println("billing.account updated", data.PublicId, data.Name, data.Password,
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
		fmt.Println("error delete accoint in billing:", err.Error())
	} else {
		fmt.Println("billing.account deleted", data.PublicId)
	}
}

func (k *Kafka) createTask(payload interface{}, service *service.Service) {
	var data models.Task
	err := k.readPayload(payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// created_at при сохранении игнорирую, оно записывается как DATETIME DEFAULT CURRENT_TIMESTAMP
	// - проекция дат создания в сервисах может пригодиться
	_, err = service.Tasker.CreateTask(data)
	if err != nil {
		fmt.Println("error create task in billing:", err.Error())
	} else {
		fmt.Println("billing.task created", data.PublicId)
	}

	// на момент регистрации этого события, таск может быть не прикреплен ни к какому пользователю
	// поэтому оставим поле пустым
	// поле Price также заполним другим BE-событием, вместе с AccountId
	err = service.SaveOrUpdateTransaction(models.Bill{
		PublicId:          uuid.New(),
		TaskId:            data.PublicId,
		TransactionReason: models.TRANSACTION_REASON_BIRD_IN_CAGE,
	})
	if err != nil {
		fmt.Println("error create bill transaction in billing:", err.Error())
	} else {
		fmt.Println("billing.task transaction created")
	}
}

// assigne task handler
func (k *Kafka) birdCageTask(payload interface{}, service *service.Service) {
	var data models.BirdCageTask
	err := k.readPayload(payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = service.Tasker.BirdCageTask(data.PublicId, data.AccountId)
	if err != nil {
		fmt.Println("error assigne task in billing:", err.Error())
	} else {
		fmt.Println("billing.task assigned", data.PublicId)
	}

	err = service.SaveOrUpdateTransaction(models.Bill{
		PublicId:  uuid.New(),
		AccountId: data.AccountId,
		Price:     rand.Intn(ZERO+MIN_PRICE) + MIN_PRICE,
	})
	if err != nil {
		fmt.Println("error create bill transaction in billing:", err.Error())
	} else {
		fmt.Println("billing.task transaction created")
	}
}

// complete task handler
func (k *Kafka) milletBowlTask(payload interface{}, service *service.Service) {
	var data models.MilletBowlTask
	err := k.readPayload(payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = service.Tasker.MilletBowlTask(data.PublicId)
	if err != nil {
		fmt.Println("error complete task in billing:", err.Error())
	} else {
		fmt.Println("billing.task completed", data.PublicId)
	}

	err = service.SaveOrUpdateTransaction(models.Bill{
		PublicId: uuid.New(),
		Price:    rand.Intn(MAX_PRICE-MIN_PRICE) + MIN_PRICE,
	})
	if err != nil {
		fmt.Println("error create bill transaction in billing:", err.Error())
	} else {
		fmt.Println("billing.task transaction created")
	}
}

func (k *Kafka) closeBillingCycle(payload interface{}, service *service.Service) {
	// проверить в таблице выплат payment, не провели ли сегодня выплату (защита от повторной обработки)
	// пройти по каждому аккаунту, посчитать заработок, обнулить
	accounts, err := service.Authorizer.GetEmployeeAccounts()
	if err != nil {
		fmt.Println("error getting employee accounts in billing:", err.Error())
		return
	}

	for _, account := range accounts {
		_ = account
		// подсчет зп за день, запись в таблицу payment - сколько мы должны
	}
}

func (k *Kafka) pay(payload interface{}, service *service.Service) {
	// выплата - отметка в payment что выплачено, можно отправку уведмления
}
