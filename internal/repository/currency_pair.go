package repository

import (
	"context"
	"currency_api/internal/models"
	"currency_api/internal/storage/postgres"
	"currency_api/pkg/log"
	"github.com/jmoiron/sqlx"
)

const defaultWell = 1

type CurrencyPairRepository struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewCurrencyPairRepository(db *sqlx.DB, logger log.Logger) *CurrencyPairRepository {
	return &CurrencyPairRepository{db: db, logger: logger}
}

func (r *CurrencyPairRepository) Create(ctx context.Context, input *models.CurrencyPairCreateInput) (*models.CurrencyPair, error) {

	ctx, cancel := context.WithTimeout(ctx, postgres.Timeout)
	defer cancel()

	query := `
		INSERT INTO "currency_pair" (currency_from, currency_to, well) 
		VALUES ($1, $2, $3)
		RETURNING currency_from, currency_to, well, updated_at
	`

	rows, err := r.db.QueryxContext(ctx, query, input.CurrencyFrom, input.CurrencyTo, defaultWell)
	if err != nil {
		// TODO: Add errors handling
		r.logger.Errorf("received the following error from postgres while create currency pair: %v", err.Error())
		return nil, err
	}
	defer rows.Close()

	newPair := &models.CurrencyPair{}
	for rows.Next() {
		if err = rows.StructScan(newPair); err != nil {
			// TODO: Add errors handling
			r.logger.Errorf("received the following error while scanning create currency pair: %v", err.Error())
			return nil, err
		}
	}

	return newPair, nil
}

func (r *CurrencyPairRepository) Get(ctx context.Context, f, t string) (*models.CurrencyPair, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.Timeout)
	defer cancel()

	// TODO: Delete "*"
	query := `SELECT * FROM currency_pair WHERE currency_from=$1 AND currency_to=$2`

	pair := models.CurrencyPair{}
	if err := r.db.GetContext(ctx, &pair, query, f, t); err != nil {
		// TODO: Add errors handling
		r.logger.Errorf("received the following error from postgres while retrieving pair: %v", err.Error())
		return nil, err
	}

	return &pair, nil
}

func (r *CurrencyPairRepository) List(ctx context.Context) (models.CurrencyPairs, error) {

	ctx, cancel := context.WithTimeout(ctx, postgres.Timeout)
	defer cancel()

	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair`

	var pairs models.CurrencyPairs
	err := r.db.SelectContext(ctx, &pairs, query)
	if err != nil {
		// TODO: Add errors handling
		r.logger.Errorf("received the following error from list currency pairs call: %v", err.Error())
		return nil, err
	}

	return pairs, nil
}

func (r *CurrencyPairRepository) UpdateCurrencyWell(ctx context.Context, exchangeInfo *models.CurrencyExchangeInfo) error {

	// FIXME: Организовать обновление в одной транзакции?
	// Или вообще обновление всей таблицы в одной транзакции? Тогда заблочим таблицу?

	r.logger.Info("UPDATING PROCESS ...")

	return nil
}
