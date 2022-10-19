package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository/postgres"
)

//go:generate minimock -i CurrencyPair -o ./mock/ -s .go -g

type CurrencyPair interface {
	Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (*models.CurrencyPair, error)
	Get(ctx context.Context, f, t string) (*models.CurrencyPair, error)
	List(ctx context.Context) (models.CurrencyPairs, error)
	UpdatePair(ctx context.Context, pair *models.CurrencyPair) error
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Pair: postgres.New(db),
	}
}

type Repository struct {
	Pair CurrencyPair
}
