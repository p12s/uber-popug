package service

import (
	"github.com/google/uuid"
	"github.com/p12s/uber-popug/task/pkg/models"
	"github.com/p12s/uber-popug/task/pkg/repository"
)

type Tasker interface {
	CreateTask(task models.Task) (int, error)
	GetTaskById(taskId int) (models.Task, error)
	GetAllTasksByAssignedAccountId(assignedAccountId int) ([]models.Task, error)
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

func (s *TaskService) GetTaskById(taskId int) (models.Task, error) {
	return s.repo.GetTaskById(taskId)
}

func (s *TaskService) GetAllTasksByAssignedAccountId(assignedAccountId int) ([]models.Task, error) {
	return s.repo.GetAllTasksByAssignedAccountId(assignedAccountId)
}

func (s *TaskService) BirdCageTask(taskId uuid.UUID, accountId uuid.UUID) error {
	return s.repo.BirdCageTask(taskId, accountId)
}

func (s *TaskService) MilletBowlTask(taskId uuid.UUID) error {
	return s.repo.MilletBowlTask(taskId)
}
