package service

import (
	"context"
	"currency_api/internal/app/currency/models"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/repository/mock"
	"currency_api/pkg/exchange_rates"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"reflect"
	"syreclabs.com/go/faker"
	"testing"
	"time"
)

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

func TestService_Create(t *testing.T) {
	mockController := minimock.NewController(t)
	defer mockController.Finish()

	ctx := context.Background()
	currencyPairMockRepo := mock.NewCurrencyPairMock(mockController)
	currencyPairCreateInput := NewCurrencyPairCreateInput(t)
	currencyPair := NewCurrencyPair(t)

	type fields struct {
		repository          *repository.Repository
		exchangeRatesClient *exchange_rates.Client
	}
	type args struct {
		ctx             context.Context
		pairCreateInput *models.CurrencyPairCreateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.CurrencyPair
		wantErr error
	}{
		{
			name: "OK",
			fields: func() fields {

				currencyPairMockRepo.
					CreateMock.
					Expect(ctx, currencyPairCreateInput).
					Return(currencyPair, nil)

				return fields{
					repository: &repository.Repository{
						Pair: currencyPairMockRepo,
					},
					exchangeRatesClient: nil,
				}
			}(),
			args: args{
				ctx:             ctx,
				pairCreateInput: currencyPairCreateInput,
			},
			want:    currencyPair,
			wantErr: nil,
		},
		// TODO: Не понимаю почему если запускать оба теста want: из первого получает значение из второго.
		/*{
			name: "ERR",
			fields: func() fields {

				currencyPairMockRepo.
					CreateMock.
					//Expect(ctx, currencyPairCreateInput).
					Return(nil, nil)

				return fields{
					repository: &repository.Repository{
						Pair: currencyPairMockRepo,
					},
					exchangeRatesClient: nil,
				}
			}(),
			args: args{
				ctx:             ctx,
				pairCreateInput: currencyPairCreateInput,
			},
			want:    nil,
			wantErr: nil,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repository:          tt.fields.repository,
				exchangeRatesClient: tt.fields.exchangeRatesClient,
			}
			got, err := s.Create(tt.args.ctx, tt.args.pairCreateInput)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
