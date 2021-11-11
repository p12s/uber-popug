package service

import (
	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/repository"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
	GetTaskById(taskId int) (models.Task, error)
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

func (s *TaskService) GetTaskById(taskId int) (models.Task, error) {
	return s.repo.GetTaskById(taskId)
}
