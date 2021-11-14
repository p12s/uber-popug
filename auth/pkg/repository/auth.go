package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/auth/pkg/models"
)

// Authorization - signup/signin
type Authorization interface {
	CreateAccount(account models.SignUpInput) (int, error)
	GetAccount(username, password string) (models.AuthAccount, error)
	GetAccountById(accountId int) (models.AuthAccount, error)
}

// Auth - repo
type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateAccount(account models.SignUpInput) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash)
		values ($1, $2, $3) RETURNING id`, accountTable)
	row := r.db.QueryRow(query, account.Name, account.Username, account.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create account: %w", err)
	}

	return id, nil
}

func (r *Auth) GetAccount(username, password string) (models.AuthAccount, error) {
	var account models.AuthAccount

	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, accountTable)
	err := r.db.Get(&account, query, username, password)
	if err != nil {
		return account, fmt.Errorf("get account: %w", err)
	}

	return account, err
}

func (r *Auth) GetAccountById(accountId int) (models.AuthAccount, error) {
	var account models.AuthAccount

	query := fmt.Sprintf(`SELECT id, name, username FROM %s WHERE id=$1`, accountTable)
	err := r.db.Get(&account, query, accountId)
	if err != nil {
		return account, fmt.Errorf("get account by id: %w", err)
	}

	return account, err
}
