package service

import (
	"context"

	"github.com/Vladroon22/TestTask-Bank-Operation/internal/entity"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/repository"
)

type Servicer interface {
	IncreaseUserBalance(c context.Context, userID int, amount float64) error
	TransferMoney(c context.Context, userFrom, userTo string, fromUserID, toUserID int, amount float64) error
	GetLastTxs(c context.Context, userID int) ([]entity.Tx, error)
}

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Servicer {
	return &Service{repo: repo}
}

func (s *Service) IncreaseUserBalance(c context.Context, userID int, amount float64) error {
	return s.repo.IncreaseUserBalance(c, userID, amount)
}

func (s *Service) TransferMoney(c context.Context, userFrom, userTo string, fromUserID, toUserID int, amount float64) error {
	return s.repo.TransferMoney(c, userFrom, userTo, fromUserID, toUserID, amount)
}

func (s *Service) GetLastTxs(c context.Context, userID int) ([]entity.Tx, error) {
	return s.repo.GetLastTxs(c, userID)
}
