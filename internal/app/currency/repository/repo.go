package repository

import (
	"context"
	"currency_api/pkg/exchange_rates"

	"github.com/jmoiron/sqlx"

	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository/postgres"
	"currency_api/pkg/log"
)

type Repository struct {
	Pair   CurrencyPair
	Logger log.Logger
}

type CurrencyPair interface {
	Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (*models.CurrencyPair, error)
	Get(ctx context.Context, f, t string) (*models.CurrencyPair, error)
	List(ctx context.Context) (models.CurrencyPairs, error)
	UpdateCurrencyWell(ctx context.Context, exchangeInfo *exchange_rates.ExchangeRatesInfo) error
}

func New(db *sqlx.DB, logger log.Logger) *Repository {
	return &Repository{
		Pair:   postgres.New(db, logger),
		Logger: logger,
	}
}
