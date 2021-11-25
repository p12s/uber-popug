package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/billing/pkg/models"
)

type Biller interface {
	SaveOrUpdateTransaction(bill models.Bill) error
}

type Bill struct {
	db *sqlx.DB
}

func NewBill(db *sqlx.DB) *Bill {
	return &Bill{db: db}
}

func (r *Task) SaveOrUpdateTransaction(bill models.Bill) error {
	// TODO переделать INSERT на save or update
	query := fmt.Sprintf(`INSERT INTO %s (public_id, account_id, task_id, transaction_reason, price)
		values ($1, $2, $3, $4, $5)`, billingTable)
	_, err := r.db.Exec(query, bill.PublicId, bill.AccountId, bill.TaskId, bill.TransactionReason, bill.Price)
	return err
}
