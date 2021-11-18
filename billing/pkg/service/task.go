package service

import (
	"github.com/p12s/uber-popug/billing/pkg/models"
	"github.com/p12s/uber-popug/billing/pkg/repository"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
}

// TaskService - service
type TaskService struct {
	repo repository.Tasker
}

// NewTaskService - constructor
func NewTaskService(repo repository.Tasker) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task models.Task) (int, error) {
	return s.repo.CreateTask(task)
}
