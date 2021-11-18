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
	EVENT_ACCOUNT_CREATED EventType = "auth.account_created"
	EVENT_ACCOUNT_UPDATED EventType = "auth.account_updated"
	EVENT_ACCOUNT_REMOVED EventType = "auth.account_removed"
)

type Event struct {
	Type  EventType
	Value interface{}
}
