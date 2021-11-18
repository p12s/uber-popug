package models

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	Employee Role = iota
	Manager
	Accountant
	Admin
)

type Account struct {
	Id        int        `json:"id" db:"id"`
	PublicId  uuid.UUID  `json:"public_id" db:"public_id" binding:"required"`
	Name      string     `json:"name" db:"name" binding:"required"`
	Username  string     `json:"username" db:"username" binding:"required"`
	Token     string     `json:"token" db:"token"`
	Role      Role       `json:"role" db:"role" binding:"required"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type TaskStatus int

const (
	Default TaskStatus = iota
	Assigned
	Completed
)

type Task struct {
	Id                int        `json:"id" db:"id"`
	PublicId          uuid.UUID  `json:"public_id" db:"public_id"`
	AssignedAccountId int        `json:"assigned_account_id" db:"assigned_account_id"`
	Description       string     `json:"description" db:"description" binding:"required"`
	Status            TaskStatus `json:"status" db:"status"`
	CreatedAt         *time.Time `json:"created_at" db:"created_at"`
}

type EventType string

const (
	EVENT_TASK_CREATED   EventType = "task.created"
	EVENT_TASK_ASSIGNED  EventType = "task.assigned"
	EVENT_TASK_COMPLETED EventType = "task.completed"
)

type Event struct {
	Type  EventType
	Value interface{}
}
