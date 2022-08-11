package repository_interfaces

import (
	"context"
	"currency_api/internal/models"
)

// FIXME: Отдельный файл для интерфейса или в файле с реализацией?

type CurrencyPairRepositoryI interface {
	Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (*models.CurrencyPair, error)
	Get(ctx context.Context, f, t string) (*models.CurrencyPair, error)
	List(ctx context.Context) (models.CurrencyPairs, error)

	UpdateCurrencyWell(ctx context.Context, exchangeInfo *models.CurrencyExchangeInfo) error
}
