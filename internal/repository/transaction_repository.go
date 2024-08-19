package repository

import (
	"context"
	"sj/internal/db/sqlc"
)

type TransactionRepository interface {
	GetAllByUserID(id string) ([]sqlc.Transaction, error)
}

type transactionRepository struct {
	db *sqlc.Queries
}

func NewTransactionRepository(db *sqlc.Queries) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetAllByUserID(id string) ([]sqlc.Transaction, error) {
	return r.db.GetAllByUserID(context.Background(), id)
}
