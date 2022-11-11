package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/transaction-service/internal/domain"
	"time"
)

type TransactionRepo struct {
	conn *sqlx.DB
}

func (c *TransactionRepo) FindAll(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	statement := "select * from transactions order by id asc"
	err := c.conn.SelectContext(ctx, &transactions, statement)
	return transactions, err
}

func (c *TransactionRepo) FindByID(ctx context.Context, transactionId int) (domain.Transaction, error) {
	statement := "select * from transactions where id = $1"
	var transaction domain.Transaction
	err := c.conn.GetContext(ctx, &transaction, statement, transactionId)
	return transaction, err
}

func (c *TransactionRepo) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	statement := "insert into transactions (sender_id, receiver_id, amount, status, updated_at) " +
		"values ($1, $2, $3, $4, $5) returning id"
	var id int64
	err := c.conn.QueryRowxContext(ctx, statement, transaction.SenderId,
		transaction.ReceiverId, transaction.Amount, transaction.Status, transaction.UpdatedAt).Scan(&id)
	transaction.Id = id
	return transaction, err
}

func (c *TransactionRepo) UpdateStatus(ctx context.Context, transactionId int64, status string) error {
	statement := "update transactions set status = $1, updated_at = $2 where id = $3"
	_, err := c.conn.ExecContext(ctx, statement, status, time.Now(), transactionId)
	return err
}

func (c *TransactionRepo) FindByStatusAndId(ctx context.Context, status string, id int64) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	statement := "select * from transactions where sender_id = $1 and status = $2 order by updated_at asc"
	err := c.conn.SelectContext(ctx, &transactions, statement, id, status)
	return transactions, err
}

func NewTransactionRepo(conn *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{
		conn: conn,
	}
}
