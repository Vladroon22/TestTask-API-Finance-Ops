package service

import (
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/entity"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/repository"
)

type Servicer interface {
	TopUpBalance(userID int, amount float64) error
	TransferMoney(fromUserID, toUserID int, amount float64) error
	GetLastTransactions(userID int) ([]entity.Transaction, error)
}

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Servicer {
	return &Service{repo: repo}
}

func (s *Service) TopUpBalance(userID int, amount float64) error {
	return s.repo.TopUpBalance(userID, amount)
}

func (s *Service) TransferMoney(fromUserID, toUserID int, amount float64) error {
	return s.repo.TransferMoney(fromUserID, toUserID, amount)
}

func (s *Service) GetLastTransactions(userID int) ([]entity.Transaction, error) {
	return s.repo.GetLastTransactions(userID)
}
