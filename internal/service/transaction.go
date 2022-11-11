package service

import (
	"context"
	"fmt"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"github.com/paw1a/transaction-service/internal/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type TransactionService struct {
	repo          repository.Transactions
	clientService Clients
}

func (t *TransactionService) FindAll(ctx context.Context) ([]domain.Transaction, error) {
	return t.repo.FindAll(ctx)
}

func (t *TransactionService) FindByID(ctx context.Context, transactionId int) (domain.Transaction, error) {
	return t.repo.FindByID(ctx, transactionId)
}

func (t *TransactionService) FindByStatusAndId(ctx context.Context, status string, id int64) ([]domain.Transaction, error) {
	return t.repo.FindByStatusAndId(ctx, status, id)
}

func (t *TransactionService) Create(ctx context.Context, transactionDto dto.CreateTransactionDto) (domain.Transaction, error) {
	if transactionDto.SenderId == transactionDto.ReceiverId {
		return domain.Transaction{}, fmt.Errorf("transaction sender and receiver are the same: %d",
			transactionDto.SenderId)
	}

	if transactionDto.Amount <= 0 {
		return domain.Transaction{}, fmt.Errorf("transaction amount must be positive value")
	}

	transaction, err := t.repo.Create(ctx, domain.Transaction{
		SenderId:   transactionDto.SenderId,
		ReceiverId: transactionDto.ReceiverId,
		Amount:     transactionDto.Amount,
		UpdatedAt:  time.Now(),
		Status:     "created",
	})

	if err != nil {
		return domain.Transaction{}, err
	}

	t.clientService.GetClientQueues()[transactionDto.SenderId] <- transaction

	return transaction, err
}

func (t *TransactionService) UpdateStatus(ctx context.Context, transactionId int64, status string) error {
	return t.repo.UpdateStatus(ctx, transactionId, status)
}

func (t *TransactionService) processTransaction(transaction domain.Transaction) error {
	err := t.clientService.Transfer(transaction.SenderId, transaction.ReceiverId, transaction.Amount)
	if err != nil {
		log.Printf("failed to make transaction with id = %d: %v", transaction.Id, err)
		err = t.UpdateStatus(context.Background(), transaction.Id, "blocked")
	} else {
		log.Printf("success transaction with id = %d", transaction.Id)
		err = t.UpdateStatus(context.Background(), transaction.Id, "done")
	}

	return err
}

func (t *TransactionService) processClientQueue(queue chan domain.Transaction) {
	for {
		select {
		case transaction := <-queue:
			err := t.processTransaction(transaction)
			if err != nil {
				return
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func NewTransactionService(repo repository.Transactions, clientService Clients) (*TransactionService, error) {
	clients, err := clientService.FindAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to find clients: %w", err)
	}

	transactionService := &TransactionService{
		repo:          repo,
		clientService: clientService,
	}

	clientQueues := transactionService.clientService.GetClientQueues()

	for _, client := range clients {
		id := client.Id
		clientQueues[id] = make(chan domain.Transaction, 64)
		createdTransactions, err := transactionService.FindByStatusAndId(
			context.Background(), "created", id)
		if err != nil {
			return nil, fmt.Errorf("failed to get client transactions with id: %d: %w",
				client.Id, err)
		}

		log.Printf("client %d", id)
		log.Printf("transactions %v", createdTransactions)

		for _, transaction := range createdTransactions {
			clientQueues[id] <- transaction
		}

		go transactionService.processClientQueue(clientQueues[id])
	}

	return transactionService, nil
}
