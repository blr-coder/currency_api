package postgres

import (
	"context"

	// DB driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"currency_api/internal/app/currency/models"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
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
		return nil, err
	}
	defer rows.Close()

	newPair := &models.CurrencyPair{}
	for rows.Next() {
		if err = rows.StructScan(newPair); err != nil {
			return nil, err
		}
	}

	return newPair, rows.Err()
}

func (r *Repository) Get(ctx context.Context, f, t string) (*models.CurrencyPair, error) {
	// TODO: Delete "*"
	query := `SELECT * FROM currency_pair WHERE currency_from=$1 AND currency_to=$2`

	pair := models.CurrencyPair{}
	if err := r.db.GetContext(ctx, &pair, query, f, t); err != nil {
		return nil, err
	}

	return &pair, nil
}

func (r *Repository) List(ctx context.Context) (models.CurrencyPairs, error) {
	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair`

	var pairs models.CurrencyPairs
	err := r.db.SelectContext(ctx, &pairs, query)
	if err != nil {
		// TODO: Add errors handling
		return nil, err
	}

	return pairs, nil
}

func (r *Repository) UpdateCurrencyWell(ctx context.Context, exchangeInfo *models.CurrencyExchangeInfo) error {

	// FIXME: Организовать обновление в одной транзакции?
	// Или вообще обновление всей таблицы в одной транзакции? Тогда заблочим таблицу?

	return nil
}
