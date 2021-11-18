package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/billing/pkg/models"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
}

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db: db}
}

func (r *Task) CreateTask(task models.Task) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (public_id, assigned_account_id, description, status)
		values ($1, $2, $3, $4) RETURNING id`, taskTable)
	row := r.db.QueryRow(query, task.PublicId.String(), task.AssignedAccountId, task.Description, models.TASK_BIRD_IN_CAGE)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create task: %w", err)
	}

	return id, nil
}
