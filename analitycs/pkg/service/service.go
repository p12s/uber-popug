package service

import (
	"github.com/p12s/uber-popug/analitycs/pkg/repository"
)

type Service struct {
	Authorizer
	Tasker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizer: NewAuthService(repos.Authorizer),
		Tasker:     NewTaskService(repos.Tasker),
	}
}
