package service

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"github.com/paw1a/transaction-service/internal/repository"
)

type ClientService struct {
	repo repository.Clients
}

func (c *ClientService) FindAll(ctx context.Context) ([]domain.Client, error) {
	return nil, nil
}

func (c *ClientService) FindByID(ctx context.Context, clientId int) (domain.Client, error) {
	return domain.Client{}, nil
}

func (c *ClientService) Create(ctx context.Context, clientDto dto.CreateClientDto) (domain.Client, error) {
	return domain.Client{}, nil
}

func (c *ClientService) Delete(ctx context.Context, clientId int) error {
	return nil
}

func NewClientService(repo repository.Clients) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
