package repository

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
)

type Clients interface {
	FindAll(ctx context.Context) ([]domain.Client, error)
	FindByID(ctx context.Context, clientId int) (domain.Client, error)
	Create(ctx context.Context, client domain.Client) (domain.Client, error)
	Delete(ctx context.Context, clientId int) error
	Transfer(senderId int64, receiverId int64, amount int64) error
}

type Transactions interface {
	FindAll(ctx context.Context) ([]domain.Transaction, error)
	FindByID(ctx context.Context, transactionId int) (domain.Transaction, error)
	FindByStatusAndId(ctx context.Context, status string, id int64) ([]domain.Transaction, error)
	Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	UpdateStatus(ctx context.Context, transactionId int64, status string) error
}
