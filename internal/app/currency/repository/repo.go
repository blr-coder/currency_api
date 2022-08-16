package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository/postgres"
)

type Repository struct {
	Pair CurrencyPair
}

type CurrencyPair interface {
	Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (*models.CurrencyPair, error)
	Get(ctx context.Context, f, t string) (*models.CurrencyPair, error)
	List(ctx context.Context) (models.CurrencyPairs, error)
	UpdateCurrencyWell(ctx context.Context, exchangeInfo *models.CurrencyExchangeInfo) error
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Pair: postgres.New(db),
	}
}
