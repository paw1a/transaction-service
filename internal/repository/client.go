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
	return nil, nil
}

func (c *ClientRepo) FindByID(ctx context.Context, clientId int) (domain.Client, error) {
	return domain.Client{}, nil
}

func (c *ClientRepo) Create(ctx context.Context, user domain.Client) (domain.Client, error) {
	return domain.Client{}, nil
}

func (c *ClientRepo) Delete(ctx context.Context, clientId int) error {
	return nil
}

func NewClientRepo(conn *sqlx.DB) *ClientRepo {
	return &ClientRepo{
		conn: conn,
	}
}
