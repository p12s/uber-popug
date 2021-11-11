package models

type TaskAccount struct {
	Id       int    `json:"id" db:"id"`
	PublicId int    `json:"public_id" db:"public_id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Token    string `json:"token" db:"token"` // чтобы не зависеть от аккаунта наврено. но как детальенне?
}

// TODO куда тут роли вставить?

type TaskStatus int

const (
	Default TaskStatus = iota
	Assigned
	Completed
)

type Task struct {
	Id                int        `json:"id" db:"id"`
	AssignedAccountId int        `json:"assigned_account_id" db:"assigned_account_id" binding:"required"`
	Description       string     `json:"description" db:"description" binding:"required"`
	Status            TaskStatus `json:"status" binding:"required"`
}
