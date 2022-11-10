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
	return c.repo.FindAll(ctx)
}

func (c *ClientService) FindByID(ctx context.Context, clientId int) (domain.Client, error) {
	return c.repo.FindByID(ctx, clientId)
}

func (c *ClientService) Create(ctx context.Context, clientDto dto.CreateClientDto) (domain.Client, error) {
	return c.repo.Create(ctx, domain.Client{
		Name:    clientDto.Name,
		Balance: clientDto.Balance,
	})
}

func (c *ClientService) Delete(ctx context.Context, clientId int) error {
	return c.repo.Delete(ctx, clientId)
}

func NewClientService(repo repository.Clients) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
