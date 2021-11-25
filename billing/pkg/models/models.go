package models

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	Employee Role = iota // TODO перевести все константы в верхний регистр
	Manager
	Accountant
	Admin
)

type Account struct {
	Id        int        `json:"id" db:"id"`
	PublicId  uuid.UUID  `json:"public_id" db:"public_id"`
	Name      string     `json:"name" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password"`
	Token     string     `json:"token"`
	Role      Role       `json:"role" db:"role"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EventType string

const (
	EVENT_ACCOUNT_CREATED EventType = "auth.created"
	EVENT_ACCOUNT_UPDATED EventType = "auth.updated"
	EVENT_ACCOUNT_REMOVED EventType = "auth.removed"

	EVENT_TASK_CREATED       EventType = "task.created"
	EVENT_TASK_BIRD_CAGED    EventType = "task.bird_caged"
	EVENT_TASK_MILLET_BOWLED EventType = "task.millet_bowled"

	EVENT_BILLING_CYCLE_CLOSED      EventType = "billing.billing_cycle_closed"
	EVENT_PAYED_TRANSACTION_APPLIED EventType = "billing.payed_transaction_applied"
)

type Event struct {
	Type  EventType
	Value interface{}
}

type TaskStatus int

const (
	TASK_DEFAULT        TaskStatus = iota
	TASK_BIRD_IN_CAGE              // Assigned
	TASK_MILLET_IN_BOWL            // Completed
)

type Task struct {
	Id                int        `json:"id" db:"id"`
	PublicId          uuid.UUID  `json:"public_id" db:"public_id"`
	AssignedAccountId int        `json:"assigned_account_id" db:"assigned_account_id"`
	Description       string     `json:"description" db:"description" binding:"required"`
	JiraId            string     `json:"jira_id,omitempty" db:"jira_id"`
	Status            TaskStatus `json:"status" db:"status"`
	CreatedAt         *time.Time `json:"created_at" db:"created_at"`
}
