package service

import (
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/analitycs/pkg/models"
	"github.com/p12s/uber-popug/analitycs/pkg/repository"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
	BirdCageTask(taskId uuid.UUID, accountId uuid.UUID) error
	MilletBowlTask(taskId uuid.UUID) error
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

func (s *TaskService) BirdCageTask(taskId uuid.UUID, accountId uuid.UUID) error {
	return s.repo.BirdCageTask(taskId, accountId)
}

func (s *TaskService) MilletBowlTask(taskId uuid.UUID) error {
	return s.repo.MilletBowlTask(taskId)
}
