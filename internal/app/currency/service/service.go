package service

import (
	"context"

	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository"
)

type Service struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Create(ctx context.Context, pair *models.CurrencyPairCreateInput) (*models.CurrencyPair, error) {
	return s.repository.Pair.Create(ctx, pair)
}

func (s Service) Get(ctx context.Context, f, t string) (*models.CurrencyPair, error) {
	return s.repository.Pair.Get(ctx, f, t)
}

func (s Service) List(ctx context.Context) (models.CurrencyPairs, error) {
	return s.repository.Pair.List(ctx)
}

func (s Service) UpdateCurrencyWell(ctx context.Context, exchangeInfo *models.CurrencyExchangeInfo) error {
	return s.repository.Pair.UpdateCurrencyWell(ctx, exchangeInfo)
}
