package repository

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
)

type Clients interface {
	FindAll(ctx context.Context) ([]domain.Client, error)
	FindByID(ctx context.Context, clientId int) (domain.Client, error)
	Create(ctx context.Context, user domain.Client) (domain.Client, error)
	Delete(ctx context.Context, clientId int) error
}
