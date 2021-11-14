package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/uber-popug/task/pkg/models"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
	GetTaskById(taskId int) (models.Task, error)
	GetAllTasksByAssignedAccountId(assignedAccountId int) ([]models.Task, error)
}

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db: db}
}

func (r *Task) CreateTask(task models.Task) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (assigned_account_id, description, status)
		values ($1, $2, $3) RETURNING id`, taskTable)
	row := r.db.QueryRow(query, task.AssignedAccountId, task.Description, models.Assigned)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create task: %w", err)
	}

	return id, nil
}

func (r *Task) GetTaskById(taskId int) (models.Task, error) {
	var task models.Task

	query := fmt.Sprintf(`SELECT id, assigned_account_id, description, status FROM %s WHERE id=$1`, taskTable)
	err := r.db.Get(&task, query, taskId)
	if err != nil {
		return task, fmt.Errorf("get task by id: %w", err)
	}

	return task, err
}

func (r *Task) GetAllTasksByAssignedAccountId(assignedAccountId int) ([]models.Task, error) {
	var tasks []models.Task

	query := fmt.Sprintf(`SELECT id, assigned_account_id, description, status FROM %s 
		WHERE assigned_account_id = $1`, taskTable)
	if err := r.db.Select(&tasks, query, assignedAccountId); err != nil {
		return tasks, fmt.Errorf("get all tasks: %w", err)
	}

	return tasks, nil
}
