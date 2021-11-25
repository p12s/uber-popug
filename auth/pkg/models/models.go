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
	PublicId  uuid.UUID  `json:"public_id" db:"public_id"`
	Name      string     `json:"name" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password,omitempty"`
	Token     string     `json:"token,omitempty"`
	Role      Role       `json:"role" db:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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

type EventType string

const (
	EVENT_ACCOUNT_CREATED       EventType = "auth.created"
	EVENT_ACCOUNT_UPDATED       EventType = "auth.updated"
	EVENT_ACCOUNT_DELETED       EventType = "auth.deleted"
	EVENT_ACCOUNT_TOKEN_UPDATED EventType = "auth.token_updated"
)

type Event struct {
	Type  EventType
	Value interface{}
}
