package models

// AuthAccount - just a account
type AuthAccount struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"-" binding:"required"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TODO add roles

type EventType string

const (
	EVENT_SIGN_UP EventType = "auth.sign_up"
)

type Event struct {
	Type  EventType // может пригодиться для разделения видов событий, но пока будет только 1
	Value interface{}
}
