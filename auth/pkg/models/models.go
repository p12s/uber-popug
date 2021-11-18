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
	Id        int        `json:"id,omitempty" db:"id"`
	PublicId  uuid.UUID  `json:"public_id" db:"public_id"`
	Name      string     `json:"name" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password,omitempty"`
	Token     string     `json:"token,omitempty"`
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
)

type Event struct {
	Type  EventType
	Value interface{}
}
