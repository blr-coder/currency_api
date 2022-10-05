package postgres

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

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

	pair := &models.CurrencyPair{}

	return pair, errors.WithStack(
		r.db.GetContext(ctx, pair, query, input.CurrencyFrom, input.CurrencyTo, 0),
	)
}

func (r *Repository) Get(ctx context.Context, f, t string) (*models.CurrencyPair, error) {

	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair 
				WHERE currency_from=$1 AND currency_to=$2`

	pair := models.CurrencyPair{}

	return &pair, errors.WithStack(r.db.GetContext(ctx, &pair, query, f, t))
}

func (r *Repository) List(ctx context.Context) (pairs models.CurrencyPairs, err error) {
	query := `SELECT currency_from, currency_to, well, updated_at FROM currency_pair`

	return pairs, errors.WithStack(r.db.SelectContext(ctx, &pairs, query))
}

func (r *Repository) UpdatePair(ctx context.Context, pair *models.CurrencyPair) error {

	logrus.Info("updatePair for:", pair)

	query := `UPDATE currency_pair SET well=$3, updated_at=$4 WHERE currency_from=$1 AND currency_to=$2`

	_, err := r.db.ExecContext(ctx, query, pair.CurrencyFrom, pair.CurrencyTo, pair.Well, pair.UpdatedAt)

	return errors.WithStack(err)
}
