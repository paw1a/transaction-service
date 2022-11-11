package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/transaction-service/internal/domain"
)

type ClientRepo struct {
	conn *sqlx.DB
}

func (c *ClientRepo) FindAll(ctx context.Context) ([]domain.Client, error) {
	var clients []domain.Client
	statement := "select * from clients order by id asc"
	err := c.conn.SelectContext(ctx, &clients, statement)
	return clients, err
}

func (c *ClientRepo) FindById(ctx context.Context, clientId int64) (domain.Client, error) {
	statement := "select * from clients where id = $1"
	var client domain.Client
	err := c.conn.GetContext(ctx, &client, statement, clientId)
	return client, err
}

func (c *ClientRepo) FindTransactionsById(ctx context.Context, clientId int64) ([]domain.Transaction, error) {
	statement := "select * from transactions where sender_id = $1 order by id asc"
	var transactions []domain.Transaction
	err := c.conn.SelectContext(ctx, &transactions, statement, clientId)
	return transactions, err
}

func (c *ClientRepo) Create(ctx context.Context, client domain.Client) (domain.Client, error) {
	statement := "insert into clients (name, balance) values ($1, $2) returning id"
	var id int64
	err := c.conn.QueryRowxContext(ctx, statement, client.Name, client.Balance).Scan(&id)
	client.Id = id
	return client, err
}

func (c *ClientRepo) Delete(ctx context.Context, clientId int64) error {
	statement := "delete from clients where id = $1"
	_, err := c.conn.ExecContext(ctx, statement, clientId)
	return err
}

func (c *ClientRepo) Transfer(senderId int64, receiverId int64, amount int64) error {
	tx, err := c.conn.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var statement string
	var balance int64
	statement = "select balance from clients where id = $1"
	tx.QueryRowContext(context.Background(), statement, senderId).Scan(&balance)
	if balance < amount {
		return fmt.Errorf("not enough money: balance = %d, amount = %d", balance, amount)
	}

	statement = "update clients set balance = $1 where id = $2"
	_, err = tx.Exec(statement, balance-amount, senderId)
	if err != nil {
		return err
	}

	statement = "select balance from clients where id = $1"
	tx.QueryRowContext(context.Background(), statement, receiverId).Scan(&balance)

	statement = "update clients set balance = $1 where id = $2"
	_, err = tx.Exec(statement, balance+amount, receiverId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func NewClientRepo(conn *sqlx.DB) *ClientRepo {
	return &ClientRepo{
		conn: conn,
	}
}
