package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/transaction-service/internal/domain"
)

type ClientRepo struct {
	conn *sqlx.DB
}

func (c *ClientRepo) FindAll(ctx context.Context) ([]domain.Client, error) {
	var clients []domain.Client
	statement := "select * from clients"
	err := c.conn.SelectContext(ctx, &clients, statement)
	return clients, err
}

func (c *ClientRepo) FindByID(ctx context.Context, clientId int) (domain.Client, error) {
	statement := "select * from clients where id = $1"
	var client domain.Client
	err := c.conn.GetContext(ctx, &client, statement, clientId)
	return client, err
}

func (c *ClientRepo) Create(ctx context.Context, client domain.Client) (domain.Client, error) {
	statement := "insert into clients (name, balance) values ($1, $2) returning id"
	var id int64
	err := c.conn.QueryRowxContext(ctx, statement, client.Name, client.Balance).Scan(&id)
	client.Id = id
	return client, err
}

func (c *ClientRepo) Delete(ctx context.Context, clientId int) error {
	statement := "delete from clients where id = $1"
	_, err := c.conn.ExecContext(ctx, statement, clientId)
	return err
}

func (c *ClientRepo) Transfer(senderId int64, receiverId int64, amount int64) error {
	return nil
}

func NewClientRepo(conn *sqlx.DB) *ClientRepo {
	return &ClientRepo{
		conn: conn,
	}
}
