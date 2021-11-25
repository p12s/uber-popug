package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/analitycs/pkg/models"
)

type Authorizer interface {
	CreateAccount(account models.Account) (int, error)
	UpdateAccount(input models.UpdateAccountInput) error
	DeleteAccountByPublicId(accountPublicId uuid.UUID) error
	GetAccount(token string) (models.Account, error)
	GetAccountById(publicId uuid.UUID) (models.Account, error)
	// это костыльный метод - опираться нужно только на publicId uuid.UUID
	GetAccountByPrimaryId(accountId int) (models.Account, error)
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

func (r *Auth) UpdateAccount(input models.UpdateAccountInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Role != nil {
		setValues = append(setValues, fmt.Sprintf("role=$%d", argId))
		args = append(args, *input.Role)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE public_id = $%d`,
		accountTable, setQuery, argId)
	args = append(args, input.PublicId.String())

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Auth) DeleteAccountByPublicId(accountPublicId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE public_id = $1`, accountTable)
	_, err := r.db.Exec(query, accountPublicId.String())
	return err
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

func (r *Auth) GetAccountByPrimaryId(accountId int) (models.Account, error) {
	var account models.Account

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, accountTable)
	err := r.db.Get(&account, query, accountId)
	if err != nil {
		return account, fmt.Errorf("get account by id: %w", err)
	}

	return account, err
}
