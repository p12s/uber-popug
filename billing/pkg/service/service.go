package service

import (
	"github.com/p12s/uber-popug/billing/pkg/repository"
)

// Service - just service
type Service struct {
	Authorization
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
