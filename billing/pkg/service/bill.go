package service

import (
	"github.com/p12s/uber-popug/billing/pkg/models"
	"github.com/p12s/uber-popug/billing/pkg/repository"
)

type Biller interface {
	SaveOrUpdateTransaction(bill models.Bill) error
}

// BillService - service
type BillService struct {
	repo repository.Biller
}

// NewBillService - constructor
func NewBillService(repo repository.Biller) *BillService {
	return &BillService{repo: repo}
}

func (s *BillService) SaveOrUpdateTransaction(bill models.Bill) error {
	return s.repo.SaveOrUpdateTransaction(bill)
}
