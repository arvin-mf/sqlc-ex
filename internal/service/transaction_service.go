package service

import (
	"sj/internal/db/sqlc"
	"sj/internal/repository"
)

type TransactionService interface {
	GetAllByUserID(id string) ([]sqlc.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{r}
}

func (s *transactionService) GetAllByUserID(id string) ([]sqlc.Transaction, error) {
	return s.repo.GetAllByUserID(id)
}
