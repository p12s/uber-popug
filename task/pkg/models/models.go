package models

import (
	"time"

	"github.com/google/uuid"
)

type Role int // TODO перевести все константы в верхний регистр

const (
	Employee Role = iota
	Manager
	Accountant
	Admin
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

type EventType string

const (
	EVENT_ACCOUNT_CREATED EventType = "auth.created"
	EVENT_ACCOUNT_UPDATED EventType = "auth.updated"
	EVENT_ACCOUNT_REMOVED EventType = "auth.removed"

	EVENT_TASK_CREATED       EventType = "task.created"
	EVENT_TASK_BIRD_CAGED    EventType = "task.bird_caged"
	EVENT_TASK_MILLET_BOWLED EventType = "task.millet_bowled"
)

type Event struct {
	Type  EventType
	Value interface{}
}
