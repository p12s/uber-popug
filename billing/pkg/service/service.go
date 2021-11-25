package service

import (
	"github.com/p12s/uber-popug/billing/pkg/repository"
)

// Service - just service
type Service struct {
	Authorizer
	Tasker
	Biller
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizer: NewAuthService(repos.Authorizer),
		Tasker:     NewTaskService(repos.Tasker),
		Biller:     NewBillService(repos.Biller),
	}
}
