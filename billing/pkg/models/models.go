package models

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	ROLE_EMPLOYEE Role = iota
	ROLE_MANAGER
	ROLE_ACCOUNTANT
	ROLE_ADMIN
)

type Account struct {
	Id        int        `json:"id,omitempty" db:"id"`
	PublicId  uuid.UUID  `json:"public_id" db:"public_id" binding:"required"`
	Name      string     `json:"name" db:"name" binding:"required"`
	Username  string     `json:"username" db:"username" binding:"required"`
	Token     string     `json:"token,omitempty" db:"token"`
	Role      Role       `json:"role" db:"role" binding:"required"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type UpdateAccountInput struct {
	PublicId uuid.UUID `json:"public_id" db:"public_id" binding:"required"`
	Name     *string   `json:"name" db:"role"`
	Password *string   `json:"password,omitempty" db:"role"`
	Role     *Role     `json:"role" db:"role"`
}

type AccountToken struct {
	PublicId uuid.UUID `json:"public_id" db:"public_id" binding:"required"`
	Token    string    `json:"token" db:"token" binding:"token"`
}

type DeleteAccountInput struct {
	PublicId uuid.UUID `json:"public_id" db:"public_id" binding:"required"`
}

type TaskStatus int

const (
	TASK_DEFAULT TaskStatus = iota
	TASK_BIRD_IN_CAGE
	TASK_MILLET_IN_BOWL
)

type Task struct {
	Id                int        `json:"id" db:"id"`
	PublicId          uuid.UUID  `json:"public_id" db:"public_id"`
	AssignedAccountId uuid.UUID  `json:"assigned_account_id" db:"assigned_account_id"`
	Description       string     `json:"description" db:"description" binding:"required"`
	JiraId            string     `json:"jira_id,omitempty" db:"jira_id"`
	Status            TaskStatus `json:"status" db:"status"`
	CreatedAt         *time.Time `json:"created_at" db:"created_at"`
}

type BirdCageTask struct {
	PublicId  uuid.UUID `json:"public_id" binding:"required"`
	AccountId uuid.UUID `json:"account_id" binding:"required"`
}

type MilletBowlTask struct {
	PublicId uuid.UUID `json:"public_id" binding:"required"`
}

type TransactionReason int

const (
	TRANSACTION_REASON_DEFAULT TransactionReason = iota
	TRANSACTION_REASON_BIRD_IN_CAGE
	TRANSACTION_REASON_MILLET_BOWL
)

type Bill struct {
	Id                int               `json:"id" db:"id"`
	PublicId          uuid.UUID         `json:"public_id" db:"public_id" binding:"required"`
	AccountId         uuid.UUID         `json:"account_id" db:"account_id"`
	TaskId            uuid.UUID         `json:"task_id" db:"task_id" binding:"required"`
	TransactionReason TransactionReason `json:"transaction_reason" db:"transaction_reason" binding:"required"`
	Price             int               `json:"price" db:"price"`
	CreatedAt         *time.Time        `json:"created_at" db:"created_at"`
}

type EventType string

const (
	EVENT_ACCOUNT_CREATED       EventType = "auth.created"
	EVENT_ACCOUNT_UPDATED       EventType = "auth.updated"
	EVENT_ACCOUNT_DELETED       EventType = "auth.deleted"
	EVENT_ACCOUNT_TOKEN_UPDATED EventType = "auth.token_updated"

	EVENT_TASK_CREATED       EventType = "task.created"
	EVENT_TASK_BIRD_CAGED    EventType = "task.bird_caged"
	EVENT_TASK_MILLET_BOWLED EventType = "task.millet_bowled"

	EVENT_BILLING_CYCLE_CLOSED              EventType = "billing.cycle_closed"
	EVENT_BILLING_PAYED_TRANSACTION_APPLIED EventType = "billing.payed_transaction_applied"
)

type Event struct {
	Type  EventType
	Value interface{}
}
