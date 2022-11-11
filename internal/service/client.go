package service

import (
	"context"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"github.com/paw1a/transaction-service/internal/repository"
	"os"
	"strconv"
	"time"
)

type ClientService struct {
	repo               repository.Clients
	transactionService Transactions
}

func (c *ClientService) FindAll(ctx context.Context) ([]domain.Client, error) {
	return c.repo.FindAll(ctx)
}

func (c *ClientService) FindById(ctx context.Context, clientId int64) (domain.Client, error) {
	return c.repo.FindById(ctx, clientId)
}

func (c *ClientService) FindTransactionsById(ctx context.Context, clientId int64) ([]domain.Transaction, error) {
	return c.repo.FindTransactionsById(ctx, clientId)
}

func (c *ClientService) Create(ctx context.Context, clientDto dto.CreateClientDto) (domain.Client, error) {
	client, err := c.repo.Create(ctx, domain.Client{
		Name:    clientDto.Name,
		Balance: clientDto.Balance,
	})
	if err != nil {
		return domain.Client{}, err
	}

	c.transactionService.ProcessTransactionQueue(client.Id)

	return client, err
}

func (c *ClientService) Delete(ctx context.Context, clientId int64) error {
	return c.repo.Delete(ctx, clientId)
}

func (c *ClientService) Transfer(senderId int64, receiverId int64, amount int64) error {
	sleepTime, _ := strconv.Atoi(os.Getenv("TRANSFER_TIME"))
	// sleep to demonstrate transaction queue
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return c.repo.Transfer(senderId, receiverId, amount)
}

func (c *ClientService) SetTransactionService(transactionService Transactions) {
	c.transactionService = transactionService
}

func NewClientService(repo repository.Clients) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
