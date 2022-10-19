package postgres

import (
	"context"
	"currency_api/internal/app/currency/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"reflect"
	"regexp"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)

// TODO: Куда то надо вынести NewMockDB наверное :)

func NewMockDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock, error) {
	t.Helper()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return sqlxDB, mock, nil
}

// TODO: Куда то надо вынести NewCurrencyPairCreateInput и NewCurrencyPair т.к. юзаем их в разных тестах

func NewCurrencyPairCreateInput(t *testing.T) *models.CurrencyPairCreateInput {
	t.Helper()
	return &models.CurrencyPairCreateInput{
		CurrencyFrom: "USD",
		CurrencyTo:   "RUB",
	}
}

func NewCurrencyPair(t *testing.T) *models.CurrencyPair {
	t.Helper()
	updatedAt := faker.Time().Backward(20 * time.Hour).UTC()
	return &models.CurrencyPair{
		CurrencyFrom: "USD",
		CurrencyTo:   "RUB",
		Well:         0,
		UpdatedAt:    &updatedAt,
	}
}

func TestRepository_Create(t *testing.T) {

	mockDB, mock, err := NewMockDB(t)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer mockDB.Close()

	currencyPair := NewCurrencyPair(t)
	currencyPairCreateInput := NewCurrencyPairCreateInput(t)

	query := `
INSERT INTO currency_pair (currency_from, currency_to, well) 
VALUES ($1, $2, $3) 
RETURNING currency_from, currency_to, well, updated_at
`

	type fields struct {
		db *sqlx.DB
	}

	type args struct {
		ctx   context.Context
		input *models.CurrencyPairCreateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.CurrencyPair
		wantErr bool
	}{
		{
			name: "OK",
			fields: func() fields {
				// TODO: Не могу разобраться что не так с ExpectPrepare ругается на запрос, что то не так  экранированием но ExpectQuery при этом отрабатывает без ошибок
				//mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(
						currencyPairCreateInput.CurrencyFrom,
						currencyPairCreateInput.CurrencyTo,
						0,
					).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"currency_from",
							"currency_to",
							"well",
							"updated_at",
						}).AddRow(
							currencyPair.CurrencyFrom,
							currencyPair.CurrencyTo,
							currencyPair.Well,
							currencyPair.UpdatedAt,
						))
				return fields{
					db: mockDB,
				}
			}(),
			args: args{
				ctx:   context.Background(),
				input: currencyPairCreateInput,
			},
			want:    currencyPair,
			wantErr: false,
		},
		{
			name: "pqError",
			fields: func() fields {
				//mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(
						currencyPairCreateInput.CurrencyFrom,
						currencyPairCreateInput.CurrencyTo,
						0,
					).
					WillReturnError(&pq.Error{})
				return fields{
					db: mockDB,
				}
			}(),
			args: args{
				ctx:   context.Background(),
				input: currencyPairCreateInput,
			},
			// TODO: ??? Почему want: nil не работает?
			want:    &models.CurrencyPair{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
