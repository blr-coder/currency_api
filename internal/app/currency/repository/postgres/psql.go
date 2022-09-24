package postgres

import (
	"context"
	"currency_api/pkg/exchange_rates"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"time"

	// DB driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"currency_api/internal/app/currency/models"
	"currency_api/pkg/log"
)

type Repository struct {
	db     *sqlx.DB
	logger log.Logger
}

func New(db *sqlx.DB, logger log.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) Create(ctx context.Context, input *models.CurrencyPairCreateInput) (*models.CurrencyPair, error) {
	query := `
		INSERT INTO currency_pair  (currency_from, currency_to, well) 
		VALUES ($1, $2, $3)
		RETURNING currency_from, currency_to, well, updated_at
	`

	rows, err := r.db.QueryxContext(ctx, query, input.CurrencyFrom, input.CurrencyTo, 0)
	if err != nil {
		r.logger.Errorf("received the following error from postgres while create pair: %v", err.Error())
		return nil, r.handleError(err)
	}
	defer rows.Close()

	newPair := &models.CurrencyPair{}
	for rows.Next() {
		if err = rows.StructScan(newPair); err != nil {
			r.logger.Errorf("received the following error while scan pair after creating: %v", err.Error())
			return nil, r.handleError(err)
		}
	}

	return newPair, rows.Err()
}

func (r *Repository) Get(ctx context.Context, f, t string) (*models.CurrencyPair, error) {

	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair 
				WHERE currency_from=$1 AND currency_to=$2`

	pair := models.CurrencyPair{}
	if err := r.db.GetContext(ctx, &pair, query, f, t); err != nil {
		r.logger.Errorf("received the following error from postgres while get pair currency_from=%s, currency_to=%s. Postgres err = %v", f, t, err.Error())
		return nil, r.handleError(err)
	}

	return &pair, nil
}

func (r *Repository) List(ctx context.Context) (models.CurrencyPairs, error) {
	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair`

	var pairs models.CurrencyPairs
	err := r.db.SelectContext(ctx, &pairs, query)
	if err != nil {
		r.logger.Errorf("received the following error from postgres while get list pairs. Postgres err = %v", err.Error())
		return nil, r.handleError(err)
	}

	return pairs, nil
}

func (r *Repository) UpdateCurrencyWell(ctx context.Context, exchangeInfo *exchange_rates.ExchangeRatesInfo) error {

	fmt.Println("UpdateCurrencyWell FOR:", exchangeInfo)

	for currency, rate := range exchangeInfo.ExchangeRates {
		now := time.Now().UTC()
		err := r.updatePair(ctx, &models.CurrencyPair{
			CurrencyFrom: exchangeInfo.Base,
			CurrencyTo:   currency,
			Well:         rate,
			UpdatedAt:    &now,
		})
		// TODO: Возможно лучше возврат массива ошибок что бы понимать на какой именно валюте что то не так?
		if err != nil {
			// TODO: Возможно можно заюзать multierr.Errors(err)
			return err
		}

		// На всякий случай
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (r *Repository) updatePair(ctx context.Context, pair *models.CurrencyPair) error {

	fmt.Println("updatePair for:", pair)

	query := `UPDATE currency_pair SET well=$3, updated_at=$4 WHERE currency_from=$1 AND currency_to=$2`

	rows, err := r.db.QueryxContext(ctx, query, pair.CurrencyFrom, pair.CurrencyTo, pair.Well, pair.UpdatedAt)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *Repository) handleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		e := NewCurrencyPairNotFound()
		return e
	}

	var pqError *pq.Error
	ok := errors.As(err, &pqError)
	if !ok {
		return NewUnexpectedBehaviorError(err.Error())
	}
	switch pqError.Code {
	case "23505":
		e := NewCurrencyPairAlreadyExist()
		e.AddParam(pqError.Table, pqError.Message)
		return e
	case "23503":
		e := NewInvalidFormError()
		e.AddParam(pqError.Column, pqError.Message)
		return e
	default:
		return NewUnexpectedBehaviorError(pqError.Message)
	}
}
