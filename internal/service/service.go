package service

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
)

type Clients interface {
	FindAll(ctx context.Context) ([]domain.Client, error)
	FindByID(ctx context.Context, clientId int) (domain.Client, error)
	Create(ctx context.Context, clientDto dto.CreateClientDto) (domain.Client, error)
	Delete(ctx context.Context, clientId int) error
}