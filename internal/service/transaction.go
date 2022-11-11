package service

import (
	"context"
	"fmt"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"github.com/paw1a/transaction-service/internal/repository"
	"time"
)

type TransactionService struct {
	repo repository.Transactions
}

func (c *TransactionService) FindAll(ctx context.Context) ([]domain.Transaction, error) {
	return c.repo.FindAll(ctx)
}

func (c *TransactionService) FindByID(ctx context.Context, transactionId int) (domain.Transaction, error) {
	return c.repo.FindByID(ctx, transactionId)
}

func (c *TransactionService) Create(ctx context.Context, transactionDto dto.CreateTransactionDto) (domain.Transaction, error) {
	if transactionDto.SenderId == transactionDto.ReceiverId {
		return domain.Transaction{}, fmt.Errorf("transaction sender and receiver are the same: %d",
			transactionDto.SenderId)
	}

	return c.repo.Create(ctx, domain.Transaction{
		SenderId:   transactionDto.SenderId,
		ReceiverId: transactionDto.ReceiverId,
		Amount:     transactionDto.Amount,
		UpdatedAt:  time.Now(),
		Status:     "created",
	})
}

func NewTransactionService(repo repository.Transactions) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}
