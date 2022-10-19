package service

import (
	"context"
	"currency_api/pkg/exchange_rates"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"time"

	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository"
)

type Service struct {
	repository          *repository.Repository
	exchangeRatesClient *exchange_rates.Client
}

func New(repository *repository.Repository, apiKey string) *Service {
	return &Service{
		repository:          repository,
		exchangeRatesClient: exchange_rates.New(apiKey),
	}
}

func (s Service) Create(ctx context.Context, pairCreateInput *models.CurrencyPairCreateInput) (*models.CurrencyPair, error) {
	return s.repository.Pair.Create(ctx, pairCreateInput)
}

func (s Service) Get(ctx context.Context, from, to string) (*models.CurrencyPair, error) {
	return s.repository.Pair.Get(ctx, from, to)
}

func (s Service) List(ctx context.Context) (models.CurrencyPairs, error) {
	return s.repository.Pair.List(ctx)
}

func (s Service) UpdateCurrencyWell(ctx context.Context, exchangeInfo *exchange_rates.ExchangeRatesInfo) error {

	var mErr error

	for currency, rate := range exchangeInfo.ExchangeRates {

		now := time.Now().UTC()
		err := s.repository.Pair.UpdatePair(ctx, &models.CurrencyPair{
			CurrencyFrom: exchangeInfo.Base,
			CurrencyTo:   currency,
			Well:         rate,
			UpdatedAt:    &now,
		})
		// TODO: Наверное не стоит обрывать выполнение цикла если какая то одна пара вернула ошибку. Возможно return []error
		if err != nil {
			//return err
			mErr = multierr.Append(mErr, err)
		}

	}

	return nil
}

func (s Service) CheckRates(ctx context.Context) error {
	listCurrencyPairs, err := s.List(ctx)
	if err != nil {
		return errors.Wrap(err, "when getting list currency pairs")
	}

	currencyMap := listCurrencyPairs.MapByCurrency()

	for from, to := range currencyMap {
		rates, err := s.exchangeRatesClient.GetRates(ctx, from, to)
		if err != nil {
			return errors.Wrapf(err, "with from state - %s, and to state - %s", from, to)
		}

		err = s.UpdateCurrencyWell(ctx, rates)
		if err != nil {
			return err
		}
	}

	return nil
}
