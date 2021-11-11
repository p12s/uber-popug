package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/task/pkg/models"
)

type Authorizer interface {
	CreateAccount(account models.TaskAccount) (int, error)
	GetAccount(token string) (models.TaskAccount, error)
	GetAccountById(accountPublicId int) (models.TaskAccount, error)
}

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateAccount(account models.TaskAccount) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (public_id, name, username, token)
		values ($1, $2, $3, $4) RETURNING id`, accountTable)
	row := r.db.QueryRow(query, account.PublicId, account.Name, account.Username, account.Token)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create account: %w", err)
	}

	return id, nil
}

func (r *Auth) GetAccount(token string) (models.TaskAccount, error) {
	var account models.TaskAccount

	query := fmt.Sprintf(`SELECT * FROM %s WHERE token=$1`, accountTable)
	err := r.db.Get(&account, query, token)
	if err != nil {
		return account, fmt.Errorf("get account: %w", err)
	}

	return account, err
}

func (r *Auth) GetAccountById(accountPublicId int) (models.TaskAccount, error) {
	var account models.TaskAccount

	query := fmt.Sprintf(`SELECT id, public_id, name, username, token FROM %s WHERE id=$1`, accountTable)
	err := r.db.Get(&account, query, accountPublicId)
	if err != nil {
		return account, fmt.Errorf("get account by id: %w", err)
	}

	return account, err
}
