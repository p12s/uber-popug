package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/task/pkg/models"
)

type Authorizer interface {
	CreateAccount(account models.Account) (int, error)
	GetAccount(token string) (models.Account, error)
	GetAccountById(publicId uuid.UUID) (models.Account, error)
}

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateAccount(account models.Account) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (public_id, name, username, token, role)
		values ($1, $2, $3, $4, $5) RETURNING id`, accountTable)
	row := r.db.QueryRow(query, account.PublicId.String(), account.Name, account.Username, account.Token, account.Role)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create account: %w", err)
	}

	return id, nil
}

func (r *Auth) GetAccount(token string) (models.Account, error) {
	var account models.Account

	query := fmt.Sprintf(`SELECT * FROM %s WHERE token=$1`, accountTable)
	err := r.db.Get(&account, query, token)
	if err != nil {
		return account, fmt.Errorf("get account: %w", err)
	}

	return account, err
}

func (r *Auth) GetAccountById(publicId uuid.UUID) (models.Account, error) {
	var account models.Account

	query := fmt.Sprintf(`SELECT * FROM %s WHERE public_id=$1`, accountTable)
	err := r.db.Get(&account, query, publicId.String())
	if err != nil {
		return account, fmt.Errorf("get account by public_id: %w", err)
	}

	return account, err
}
