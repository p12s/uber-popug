package service

import (
	"github.com/p12s/uber-popug/auth/pkg/repository"
)

// Service - just service
type Service struct {
	Authorizer
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizer: NewAuthService(repos.Authorizer),
	}
}
