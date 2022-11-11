package service

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"github.com/paw1a/transaction-service/internal/repository"
)

type ClientService struct {
	repo         repository.Clients
	clientQueues map[int64]chan domain.Transaction
}

func (c *ClientService) FindAll(ctx context.Context) ([]domain.Client, error) {
	return c.repo.FindAll(ctx)
}

func (c *ClientService) FindByID(ctx context.Context, clientId int) (domain.Client, error) {
	return c.repo.FindByID(ctx, clientId)
}

func (c *ClientService) Create(ctx context.Context, clientDto dto.CreateClientDto) (domain.Client, error) {
	client, err := c.repo.Create(ctx, domain.Client{
		Name:    clientDto.Name,
		Balance: clientDto.Balance,
	})
	if err != nil {
		return domain.Client{}, err
	}

	c.clientQueues[client.Id] = make(chan domain.Transaction, 64)

	return client, err
}

func (c *ClientService) Delete(ctx context.Context, clientId int) error {
	return c.repo.Delete(ctx, clientId)
}

func (c *ClientService) Transfer(senderId int64, receiverId int64, amount int64) error {
	return c.repo.Transfer(senderId, receiverId, amount)
}

func (c *ClientService) GetClientQueues() map[int64]chan domain.Transaction {
	return c.clientQueues
}

func NewClientService(repo repository.Clients) *ClientService {
	return &ClientService{
		repo:         repo,
		clientQueues: make(map[int64]chan domain.Transaction),
	}
}
